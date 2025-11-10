package xutils

import (
	"fmt"
	"testing"
	"time"
)

func TestGO(t *testing.T) {
	GO(func() {
		fmt.Println("------:123456")
		panic("1234456")
		fmt.Println("------:abcdef")
	})
	fmt.Println("------:111111111")
	time.Sleep(3 * time.Second)
}
