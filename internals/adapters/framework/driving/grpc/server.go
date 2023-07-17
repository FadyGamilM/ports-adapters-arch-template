package rpc

import (
	"log"
	"net"

	"github.com/FadyGamilM/hex/internals/adapters/framework/driving/grpc/pb"
	"github.com/FadyGamilM/hex/internals/ports"
	"google.golang.org/grpc"
)

type Adapter struct {
	api ports.ApiPort
	pb.UnimplementedTransactionServiceServer
}

// run method is used to start our grpc service 
func (a *Adapter)Run () {
	lis, err := net.Listen("tcp", ":5050")
	if err!= nil {
		log.Fatalf("error while creating a listener on port 5050 : %v \n", err)
	}

	gs := grpc.NewServer()

	pb.RegisterTransactionServiceServer(gs, a)
	if err:= gs.Serve(lis) ; err != nil {
		log.Fatalf("failed to serve the grpc server on port 5050 : %v \n", err)
	}
}
 

// factory for the adapter
func NewAdapter(api_layer ports.ApiPort) *Adapter{
	return &Adapter{api: api_layer}
}