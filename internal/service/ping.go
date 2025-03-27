package service

import (
	"context"

	pb "http-server/api/ping"
)

type PingService struct {
	pb.UnimplementedPingServer
}

func NewPingService() *PingService {
	return &PingService{}
}

func (s *PingService) GetPing(ctx context.Context, req *pb.GetPingRequest) (*pb.GetPingReply, error) {
	return &pb.GetPingReply{
		Msg: "pong",
	}, nil
}
