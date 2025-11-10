package database

import (
	"context"
	"crypto/tls"
	"runtime"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/redis/go-redis/v9/maintnotifications"
)

var redisConn *redis.ClusterClient

func InitRedis(c *RedisClusterConfig) error {
	return initRedis(c)
}

func initRedis(c *RedisClusterConfig) error {
	//redis options
	redisOptions := redis.ClusterOptions{
		Addrs:    c.Host,
		Username: c.Username,
		Password: c.Password,

		//连接池容量及闲置连接数量
		PoolSize:     c.PoolSize,     //链接池最大链接数，默认为cup * 5。
		MinIdleConns: c.MinIdleConns, //在启动阶段，链接池最小链接数，并长期维持idle状态的链接数不少于指定数量。
		MaxIdleConns: c.MaxIdleConns,
		//超时设置
		DialTimeout:     5 * time.Second,    //建立链接超时时间，默认为5秒。
		ReadTimeout:     3 * time.Second,    //读超时，默认3秒，-1表示取消读超时。
		WriteTimeout:    3 * time.Second,    //写超时，默认等于读超时。
		PoolTimeout:     5 * time.Second,    //当所有连接都处在繁忙状态时，客户端等待可用连接的最大等待时长，默认为读超时+1秒。
		ConnMaxLifetime: 3600 * time.Second, //链接存活时长
		//命令执行失败时的重试策略
		MaxRetries:      3,                      //命令执行失败时，最多重试多少次，默认为0即不重试。
		MinRetryBackoff: 8 * time.Microsecond,   //每次计算重试间隔时间的下限，默认8毫秒，-1表示取消间隔。
		MaxRetryBackoff: 512 * time.Microsecond, //每次计算重试间隔时间的上限，默认512毫秒，-1表示取消间隔。
		MaintNotificationsConfig: &maintnotifications.Config{
			Mode: maintnotifications.ModeDisabled,
		},

		//仅当客户端执行命令时需要从连接池获取连接时，如果连接池需要新建连接时则会调用此钩子函数。
		OnConnect: func(ctx context.Context, cn *redis.Conn) error {
			return nil
		},
	}
	//tls
	if c.TLS {
		redisOptions.MaintNotificationsConfig = &maintnotifications.Config{
			Mode: maintnotifications.ModeDisabled,
		}
		redisOptions.TLSConfig = &tls.Config{
			MinVersion: tls.VersionTLS12,
			ServerName: c.ServerName,
		}
	}

	//create redis connection
	redisConn = redis.NewClusterClient(&redisOptions)
	if err := redisConn.Ping(context.Background()).Err(); err != nil {
		return err
	}

	//Finalizer
	runtime.SetFinalizer(redisConn, func(conn *redis.ClusterClient) {
		conn.Close()
	})

	//return
	return nil
}

func CloseRD() {
	//保证传入的参数在这个方法被调用之前不被垃圾回收器回收掉
	runtime.KeepAlive(redisConn)
	if redisConn != nil {
		redisConn.Close()
	}
	redisConn = nil
}

func GetRD() *redis.ClusterClient {
	return redisConn
}
