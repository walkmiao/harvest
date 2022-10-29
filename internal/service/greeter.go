package service

import (
	"context"
	"harvest/internal/biz"

	pb "harvest/api/helloworld/v1"
)

type GreeterService struct {
	repo *biz.GreeterUsecase
	pb.UnimplementedGreeterServer
}

func NewGreeterService(ucase *biz.GreeterUsecase) *GreeterService {
	return &GreeterService{
		repo: ucase,
	}
}

func (s *GreeterService) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "hello," + req.Name + "!"}, nil
}
