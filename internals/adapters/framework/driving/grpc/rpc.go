package rpc

import (
	"context"
	"log"

	"github.com/FadyGamilM/hex/internals/adapters/framework/driving/grpc/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a Adapter) PerformDeposite(ctx context.Context, req *pb.Transaction) (*pb.CurrentBalance, error) {
	if req.GetAmount() == 0 {
		log.Printf("The request is not valid, missing required fields")
		return &pb.CurrentBalance{}, status.Error(codes.InvalidArgument, "Unexpeced error")

	}
	// call the application layer
	res, err := a.api.PerformDeposite(req.Amount)
	if err!= nil {
		log.Printf("FRAMEWORK DRIVING LAYER | PerformDeposite() | error while getting the response : %v \n", err)
		return &pb.CurrentBalance{}, status.Error(codes.InvalidArgument, "Unexpeced error")
	}
	return &pb.CurrentBalance{Balance: res}, nil
}

func (a Adapter) PerformWithdraw(ctx context.Context, req *pb.Transaction) (*pb.CurrentBalance, error) {
	if req.GetAmount() == 0 {
		log.Printf("The request is not valid, missing required fields")
		return &pb.CurrentBalance{}, status.Error(codes.InvalidArgument, "Unexpeced error")

	}
	// call the application layer
	res, err := a.api.PerformWithdraw(req.Amount)
	if err!= nil {
		log.Printf("FRAMEWORK DRIVING LAYER | PerformDeposite() | error while getting the response : %v \n", err)
		return &pb.CurrentBalance{}, status.Error(codes.InvalidArgument, "Unexpeced error")
	}
	return &pb.CurrentBalance{Balance: res}, nil
}
