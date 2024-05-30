package app_serv

import (
	v1 "be_demo/api/demo/v1"
	"be_demo/internal/biz/app_biz"
	"be_demo/internal/repository/app_repo"
	"context"
)

// GreeterService is a greeter service.
type GreeterService struct {
	v1.UnimplementedGreeterServer

	uc *app_biz.GreeterUsecase
}

// NewGreeterService new a greeter service.
func NewGreeterService(uc *app_biz.GreeterUsecase) *GreeterService {
	return &GreeterService{uc: uc}
}

// SayHello implements helloworld.GreeterServer.
func (s *GreeterService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	g, err := s.uc.CreateGreeter(ctx, &app_repo.Greeter{Hello: in.Name})
	if err != nil {
		return nil, err
	}
	return &v1.HelloReply{Message: "Hello " + g.Hello}, nil
}
