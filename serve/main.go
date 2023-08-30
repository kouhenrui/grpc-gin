package main

import (
	"flag"
	"google.golang.org/grpc"
	"grpc-gin/etcd"
	"grpc-gin/proto/auth"
	authserve "grpc-gin/serve/auth"
	"net"
)

var (
	port        = flag.String("port", ":50001", "")
	serviceName = "etcd-client"
)

func main() {

	flag.Parse()
	lsn, err := net.Listen("tcp", *port)
	if err != nil {
		panic(err)
	}
	serve := grpc.NewServer()
	auth.RegisterAuthInterServer(serve, &authserve.Service{})
	err = etcd.EtcdConnDis(serviceName, *port)
	if err != nil {
		panic(err)
	}
	if err = serve.Serve(lsn); err != nil {
		panic(err)
	}
}
