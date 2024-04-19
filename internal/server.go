package internal // github.com/bartmika/serialreader-server/internal/server.go

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/yjuw/serialreader-server/proto"
)

type SerialReaderServer struct {
	port       int
	grpcServer *grpc.Server
}

func New(port int) *SerialReaderServer {
	return &SerialReaderServer{
		port:       port,
		grpcServer: nil,
	}
}

// Function will consume the main runtime loop and run the business logic
// of the application.
func (s *SerialReaderServer) RunMainRuntimeLoop() {
	// Open a TCP server to the specified localhost and environment variable
	// specified port number.
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", s.port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Initialize our gRPC server using our TCP server.
	grpcServer := grpc.NewServer()

	// Save reference to our application state.
	s.grpcServer = grpcServer

	// For debugging purposes only.
	log.Printf("gRPC server is running.")

	// Block the main runtime loop for accepting and processing gRPC requests.
	pb.RegisterSerialReaderServer(grpcServer, &SerialReaderServerImpl{
		// DEVELOPERS NOTE:
		// We want to attach to every gRPC call the following variables...
		// ...
	})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// Function will tell the application to stop the main runtime loop when
// the process has been finished.
func (s *SerialReaderServer) StopMainRuntimeLoop() {
	log.Printf("Starting graceful shutdown now...")

	// Finish any RPC communication taking place at the moment before
	// shutting down the gRPC server.
	s.grpcServer.GracefulStop()
}
