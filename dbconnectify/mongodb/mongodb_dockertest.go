package mongodb

import (
	"context"
	"fmt"
	"strings"

	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDockerTestConf struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string

	ResourceExpired uint
	network         *docker.Network
	image           string
}

func (m *MongoDockerTestConf) URI() string {
	return fmt.Sprintf("mongodb://%s:%s@%s:%s", m.Username, m.Password, m.Host, m.Port)
}

func (m *MongoDockerTestConf) ImageVersion(network *docker.Network, version string) *dockertest.RunOptions {
	m.network = network
	m.InitConf(version)
	return &dockertest.RunOptions{
		Repository: `mongo`,
		Name:       `dockertest-mongo-` + m.network.Name,
		Tag:        m.image,
		NetworkID:  network.ID,
		Env: []string{
			`MONGO_INITDB_ROOT_USERNAME=` + m.Username,
			`MONGO_INITDB_ROOT_PASSWORD=` + m.Password,
		},
	}
}

func (m *MongoDockerTestConf) ConnectClient(resource *dockertest.Resource) (conn *mongo.Client, err error) {
	if m.ResourceExpired != 0 {
		resource.Expire(m.ResourceExpired)
	}

	hostAndPort := resource.GetHostPort("27017/tcp")
	port := strings.Split(hostAndPort, ":")[1]
	m.Host = strings.Split(hostAndPort, ":")[0]
	m.Port = port

	ctx := context.Background()
	opts := options.Client().ApplyURI(m.URI())
	conn, err = OpenConnMongoClient(ctx, opts)
	if err != nil {
		panic(err)
	}

	return
}

func (m *MongoDockerTestConf) InitConf(version string) {
	if m.Username == "" {
		m.Username = "root"
	}
	if m.Password == "" {
		m.Password = "root"
	}
	if m.Database == "" {
		m.Database = "dockertest"
	}
	if m.image == "" {
		if version != "" {
			m.image = version
		} else {
			m.image = "latest"
		}
	}
}
