package main

import (
	"../../util/netutil"
	"../etcdservice"
	"../protos/datapb"
	"flag"
	"google.golang.org/grpc/reflection"
)

func main() {
	port := flag.Int("port", 0, "port to listen")
	flag.Parse()
	l, err := netutil.ListenLocalTCP(*port)
	if err != nil {
		panic(err)
	}

	dbserver := &DatabaseServer{}
	grpcServer := etcdservice.NewServer()
	datapb.RegisterDatabaseServiceServer(grpcServer, dbserver)
	reflection.Register(grpcServer)

	addr := l.Addr().String()
	etcdservice.RegisterForever(etcdservice.DATA, addr, addr)
	if err := grpcServer.Serve(l); err != nil {
		panic(err)
	}
}
