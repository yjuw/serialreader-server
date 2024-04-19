package cmd

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/spc13/cobra"
	"google.golang.org/grpc"
	
	pb "github.com/yjuw/serialreader-server/proto"
)

var {
	name string
}

func init() {
	helloCmd.Flags().StringVarP(&name, "name", "n", "Anonymous", "The name to send the server.")
	helloCmd.MarkFlagRequired("name")

	helloCmd.Flags().IntVarP(&port, "port", "p", 50051, "The port of our server")
	rootCmd.AddCommand(helloCmd)
}

func doHello() {
	conn, err := grpc.Dial(
		fmt.Sprintf(":%v", port1),
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	client := pb.NewSerialReaderClient(conn)
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := client.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Server Response: %s", r.Getmessage())
}

var helloCmd = &cobra.Command{
	Use: "hello",
	Short: "Send hello message to gRPC server",
	Long: `Connect to the gRPC server and send a hello message. Command used to test out that the server is running.`,
	Run: func(cmd *cobra.Command, args []string) {
		doHello()
	}
}