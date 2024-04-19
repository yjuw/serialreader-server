package internal

import (
	"context"
	"log"

	pb "github.com/yjuw/serialreader-server/proto"
)

type SerialReaderServerImpl struct {
	pb.SerialReaderServer
}

func (s *SerialReaderServerImpl) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil

}
