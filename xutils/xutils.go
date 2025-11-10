package xutils

import (
	"context"
	"errors"
	"net"
	"runtime"
	"strings"
	"time"
	"unsafe"

	"github.com/tonly18/xgin/logger"
)

var Timeout = errors.New("timeout")

// 并发协程数量
var goChanNumber = make(chan struct{}, 600)

type eface struct {
	typ unsafe.Pointer
	ptr unsafe.Pointer
}

// IsNil 值判空
func IsNil(v any) bool {
	if v == nil {
		return true
	}

	ep := (*eface)(unsafe.Pointer(&v))
	if ep == nil {
		return true
	}

	return ep.typ == nil || uintptr(ep.ptr) == 0x0
}

// IsPrivateIP 判断一个 IP 是否为内网 / 本地地址
func IsPrivateIP(ips string) bool {
	ip := net.ParseIP(ips)
	if ip == nil {
		return false
	}

	// 转成 IPv4（若为 IPv6 则不变）
	ip4 := ip.To4()
	if ip4 != nil {
		// IPv4 检查
		switch {
		case ip4[0] == 10:
			return true // 10.0.0.0/8
		case ip4[0] == 172 && ip4[1] >= 16 && ip4[1] <= 31:
			return true // 172.16.0.0/12
		case ip4[0] == 192 && ip4[1] == 168:
			return true // 192.168.0.0/16
		case ip4[0] == 127:
			return true // 回环地址 127.0.0.0/8
		case ip4[0] == 169 && ip4[1] == 254:
			return true // 链路本地地址 169.254.0.0/16
		default:
			return false
		}
	}

	// IPv6 检查
	if ip.IsLoopback() || ip.IsLinkLocalUnicast() || ip.IsLinkLocalMulticast() {
		return true
	}
	// fc00::/7 -> Unique Local Address (私有 IPv6)
	if ip[0]&0xfe == 0xfc {
		return true
	}

	return false
}

// IsLocalIP 是否为本机IP
func IsLocalIP(ip string) bool {
	parsed := net.ParseIP(ip)
	if parsed == nil {
		return false
	}

	// 回环地址 (127.0.0.1 / ::1)
	if parsed.IsLoopback() {
		return true
	}

	// 检查是否为本机网卡
	addrs, _ := net.InterfaceAddrs()
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok {
			if ipnet.IP.Equal(parsed) {
				return true
			}
		}
	}

	return false
}

func GO(handler func()) error {
	timer := time.NewTimer(500 * time.Millisecond)
	defer timer.Stop()

	select {
	case <-timer.C:
		return Timeout
	case goChanNumber <- struct{}{}:
		go func() {
			defer func() {
				_ = <-goChanNumber
				if err := recover(); err != nil {
					pc, _, _, _ := runtime.Caller(2)
					fn := runtime.FuncForPC(pc).Name()
					if i := strings.LastIndex(fn, "/"); i != -1 {
						fn = fn[i+1:]
					}
					logger.Errorf(context.Background(), "GO panic recover(%v) %d/%d, error: %+v", fn, len(goChanNumber), cap(goChanNumber), err)
				}
			}()
			handler()
		}()

		return nil
	}
}

// BytesToString []byte转string
func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// StringToBytes string 转[]byte
func StringToBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(
		&struct {
			string
			Cap int
		}{s, len(s)},
	))
}
