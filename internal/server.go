package internal

import (
	"fmt"
	"log"
	"net"

	"google.golang.org//grpc"
	pb "github.com/yjuw/serialreader-server/proto"
)

type SerialReaderServer struct {
	port int
	grpcServer *grpc.Server
}

func New(port int) *SerialReaderServer {
	return &SerialReaderServer{
		port: port,
		grpcServer: nil,
	}
}

func (s *SerialReaderServer) RunMainRuntimeLoop() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", s.port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	// save reference to application state
	s.grpcServer = grpcServer

	log.Printf("gRPC server is running")

	pb.RegisterSerialReaderServer(grpcServer, &SerialReaderServerImpl {

	})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

func (s *SerialReaderServer) StopMainRuntimeLoop() {
	log.Printf("starting graceful shutdown now...")
	s.grpcServer.GracefulStop()
}

}

