package service

import (
	"context"
	"errors"

	pb "github.com/kantapapan/grpc-demo/pb"
)

// MyCatService ...
type MyCatService struct {
}

// GetMyCat ...
func (s *MyCatService) GetMyCat(ctx context.Context, message *pb.GetMyCatMessage) (*pb.MyCatResponse, error) {
	switch message.TargetCat {
	case "tama":
		//たまはメインクーン
		return &pb.MyCatResponse{
			Name: "tama",
			Kind: "mainecoon",
		}, nil
	case "mike":
		//ミケはノルウェージャンフォレストキャット
		return &pb.MyCatResponse{
			Name: "mike",
			Kind: "Norwegian Forest Cat",
		}, nil
	}
	return nil, errors.New("Not Found YourCat")
}
