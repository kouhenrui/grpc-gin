package auth

import (
	"context"
	protoauth "grpc-gin/proto/auth"
)

/**
 * @ClassName auth
 * @Description TODO
 * @Author khr
 * @Date 2023/8/25 16:30
 * @Version 1.0
 */

type Service struct {
	protoauth.UnimplementedAuthInterServer
}

func (Service) Login(context.Context, *protoauth.LoginReq) (*protoauth.LoginRes, error) {

	return &protoauth.LoginRes{
		Jwt: "123456",
	}, nil
}
func (Service) Info(context.Context, *protoauth.InfoReq) (*protoauth.InfoRes, error) {
	return &protoauth.InfoRes{Name: "张三，你好"}, nil
}
