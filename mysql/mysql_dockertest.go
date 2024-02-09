package mysql

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
)

type MysqlDockerTestConf struct {
	User     string
	Password string
	Host     string
	Port     int
	DB       string

	ResourceExpired uint
	network         *docker.Network
	image           string
}

func (p *MysqlDockerTestConf) DBURL() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		p.User, p.Password, p.Host, p.Port, p.DB)
}

func (p *MysqlDockerTestConf) UriWithOutDB() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/",
		p.User, p.Password, p.Host, p.Port)
}

func (p *MysqlDockerTestConf) ImageVersion(network *docker.Network, maridb bool, version string) *dockertest.RunOptions {
	p.network = network
	p.InitConf(version)

	name := "mysql"
	if maridb {
		name = "mariadb"
	}

	options := &dockertest.RunOptions{
		Name:       name + "-" + p.network.Name,
		Repository: name,
		Tag:        version,
		Env: []string{
			"MYSQL_ROOT_PASSWORD=" + p.Password,
		},
	}

	return options
}

func (p *MysqlDockerTestConf) ConnectSqlx(resource *dockertest.Resource) (conn *sqlx.DB, err error) {
	if p.ResourceExpired != 0 {
		resource.Expire(p.ResourceExpired)
	}

	hostAndPort := resource.GetHostPort("3306/tcp")

	port, err := strconv.Atoi(strings.Split(hostAndPort, ":")[1])
	if err != nil {
		return
	}
	p.Host = strings.Split(hostAndPort, ":")[0]
	p.Port = port

	conn = MustOpenConnSqlx(p.UriWithOutDB())
	_, err = conn.Exec("CREATE DATABASE IF NOT EXISTS " + p.DB)
	conn.Close()
	if err != nil {
		panic(err)
	}

	conn = MustOpenConnSqlx(p.DBURL())
	return
}

func (p *MysqlDockerTestConf) InitConf(version string) {
	if p.User == "" {
		p.User = "root"
	}
	if p.Password == "" {
		p.Password = "root"
	}
	if p.DB == "" {
		p.DB = "dockertest"
	}
	if p.image == "" {
		if version != "" {
			p.image = version
		} else {
			p.image = "latest"
		}
	}
}
