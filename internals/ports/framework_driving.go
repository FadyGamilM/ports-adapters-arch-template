package ports

import (
	"context"

	"github.com/FadyGamilM/hex/internals/adapters/framework/driving/grpc/pb"
)

type GrpcPort interface {
	Run()
	PerformDeposite(ctx context.Context, req *pb.Transaction) (*pb.CurrentBalance, error)
	PerformWithdraw(ctx context.Context, req *pb.Transaction) (*pb.CurrentBalance, error)
}
