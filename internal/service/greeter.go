package service

import (
	"context"

	v1 "github.com/kvii/greet/api/helloworld/v1"
	"github.com/kvii/greet/internal/biz"
)

// GreeterService is a greeter service.
type GreeterService struct {
	v1.UnimplementedGreeterServer

	uc *biz.GreeterUsecase
}

// NewGreeterService new a greeter service.
func NewGreeterService(uc *biz.GreeterUsecase) *GreeterService {
	return &GreeterService{uc: uc}
}

// SayHello 方法进行一个招呼的打。
// 当用户名为 "404" 时应报错。错误原因是 "USER_NOT_FOUND"。
// 当用户名为 "400" 时应报错。错误原因是 "GREETER_UNSPECIFIED"。
func (s *GreeterService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	g, err := s.uc.CreateGreeter(ctx, &biz.Greeter{Hello: in.Name})
	if err != nil {
		return nil, err
	}
	return &v1.HelloReply{Message: "Hello " + g.Hello}, nil
}
