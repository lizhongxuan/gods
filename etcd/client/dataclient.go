package main

import (
	"../etcdservice"
	"../protos/datapb"
	"google.golang.org/grpc/keepalive"
	"time"
)

var DB datapb.DatabaseServiceClient
var clientKeepaliveParams = keepalive.ClientParameters{
	Time:                time.Second * 30,
	Timeout:             time.Second * 20,
	PermitWithoutStream: true,
}

func init() {
	//服务器的名字 etcdservice.DATA
	conn := etcdservice.NewClientConn(etcdservice.DATA)
	DB = datapb.NewDatabaseServiceClient(conn)
}
