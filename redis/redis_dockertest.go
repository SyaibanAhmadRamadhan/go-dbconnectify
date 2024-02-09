package redis

import (
	"strconv"
	"strings"

	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/redis/rueidis"
)

type RedisDockerTestConf struct {
	Host     string
	Port     int
	DB       int
	Password string

	ResourceExpired uint
	image           string
}

func (p *RedisDockerTestConf) ImageVersion(network *docker.Network, version string) *dockertest.RunOptions {
	p.InitConf(version)

	options := &dockertest.RunOptions{
		Name:       "redis-" + network.Name,
		Repository: "redis",
		Tag:        version,
		Env:        []string{},
		Cmd: []string{
			"redis-server",
			"--requirepass",
			p.Password,
		},
	}

	return options
}

func (p *RedisDockerTestConf) ConnectRueidis(resource *dockertest.Resource) (adapter *RueidisAdapter, err error) {
	if p.ResourceExpired != 0 {
		resource.Expire(p.ResourceExpired)
	}

	hostAndPort := resource.GetHostPort("6379/tcp")

	port, err := strconv.Atoi(strings.Split(hostAndPort, ":")[1])
	if err != nil {
		return
	}
	p.Host = strings.Split(hostAndPort, ":")[0]
	p.Port = port

	adapter = OpenConnRueidis(rueidis.ClientOption{
		InitAddress: []string{hostAndPort},
		Password:    p.Password,
		SelectDB:    p.DB,
	})

	return
}

func (p *RedisDockerTestConf) InitConf(version string) {
	if p.Password == "" {
		p.Password = "root"
	}
	if p.image == "" {
		if version != "" {
			p.image = version
		} else {
			p.image = "latest"
		}
	}

}
