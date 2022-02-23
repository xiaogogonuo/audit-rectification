package initialize

import (
	"fmt"
	nebula "github.com/vesoft-inc/nebula-go/v3"
	"log"
)

const (
	address  = "192.168.149.10"
	port     = 9669
	username = "root"
	password = "nebula"
)

func NebulaInit() *nebula.Session {
	hostAddress := nebula.HostAddress{Host: address, Port: port}
	hostList := []nebula.HostAddress{hostAddress}
	pool, err := nebula.NewConnectionPool(hostList, nebula.GetDefaultConf(), nebula.DefaultLogger{})
	if err != nil {
		log.Fatal(fmt.Sprintf("Fail to initialize the connection pool, host: %s, port: %d, %s", address, port, err.Error()))
	}
	session, err := pool.GetSession(username, password)
	if err != nil {
		log.Fatal(fmt.Sprintf("Fail to create a new session from connection pool, username: %s, password: %s, %s",
			username, password, err.Error()))
	}
	return session
}
