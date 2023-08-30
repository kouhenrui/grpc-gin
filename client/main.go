package main

import (
	"fmt"
	"grpc-gin/client/src/route"
)

func main() {
	r := route.InitRoute()
	if err := r.Run(":3000"); err != nil { //, "https/certificate.crt", "https/private.key"); err != nil {
		fmt.Println(fmt.Errorf("端口占用,err:%v\n", err))
	}
}
