package database

import (
	"context"
	"fmt"
	"testing"
)

func init() {
	rdconf := &RedisClusterConfig{
		Host: []string{"54.69.237.139:8084", "54.69.237.139:8085", "54.69.237.139:8086"},
	}
	if err := initRedis(rdconf); err != nil {
		panic(err)
	}
}

func TestRedisCluster(t *testing.T) {
	rd := GetRD()
	fmt.Printf("------rd:%+v\n", rd)

	cmd := rd.Get(context.Background(), "animal")
	res, err := cmd.Result()
	fmt.Printf("------err:%+v\n", err)
	fmt.Printf("------res:%+v\n", res)

}
