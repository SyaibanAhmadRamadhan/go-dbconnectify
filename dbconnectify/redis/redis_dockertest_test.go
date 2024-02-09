package redis

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/ory/dockertest/v3"
	"github.com/redis/rueidis"

	"github.com/SyaibanAhmadRamadhan/gocatch/gcommon"
	"github.com/SyaibanAhmadRamadhan/gocatch/ginfra"
)

func TestRedisDockerTest(t *testing.T) {
	dockerTest := ginfra.InitDockerTest()
	defer dockerTest.CleanUp()

	redisDockerTestConf := RedisDockerTestConf{}

	var client rueidis.Client
	var adapter *RueidisAdapter

	dockerTest.NewContainer(redisDockerTestConf.ImageVersion(dockerTest, ""), func(res *dockertest.Resource) error {
		time.Sleep(2 * time.Second)
		conn, err := redisDockerTestConf.ConnectRueidis(res)
		client = conn.Conn
		adapter = conn
		gcommon.PanicIfError(err)

		return nil
	})

	ctx := context.Background()
	// SET key val NX
	err := adapter.Hset(ctx, "asd", "field1", "asd").Error()
	gcommon.PanicIfError(err)
	errStr, _ := adapter.Hsetnx(ctx, "asd", "field1", "asd").ToInt64()
	fmt.Println(errStr)
	// HGETALL hm
	key, err := client.Do(ctx, client.B().Hgetall().Key("asd").Build()).ToMap()
	gcommon.PanicIfError(err)
	fmt.Println(key)

}
