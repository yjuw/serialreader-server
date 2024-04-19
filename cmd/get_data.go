package cmd // github.com/bartmika/serialreader-server/cmd/get_data.go

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"

	// "google.golang.org/grpc/credentials"

	pb "github.com/yjuw/serialreader-server/proto"
)

func init() {
	// The following are optional and will have defaults placed when missing.
	getDataCmd.Flags().IntVarP(&port, "port", "p", 50051, "The port of our server.")
	rootCmd.AddCommand(getDataCmd)
}

func doGetData() {
	// Set up a direct connection to the gRPC server.
	conn, err := grpc.Dial(
		fmt.Sprintf(":%v", port),
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	// Set up our protocol buffer interface.
	client := pb.NewSerialReaderClient(conn)
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Perform our gRPC request.
	tsd, err := client.GetSparkFunWeatherShieldData(ctx, &pb.GetTimeSeriesData{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	// Print out the gRPC response.
	log.Println("Server Response:")
	fmt.Println("Status: ", tsd.Status)
	fmt.Println("HumidityValue: ", tsd.HumidityValue)
	fmt.Println("HumidityUnit: ", tsd.HumidityUnit)
	fmt.Println("TemperatureValue: ", tsd.TemperatureValue)
	fmt.Println("TemperatureUnit: ", tsd.TemperatureUnit)
	fmt.Println("PressureValue: ", tsd.PressureValue)
	fmt.Println("PressureUnit: ", tsd.PressureUnit)
	fmt.Println("TemperatureBackupValue: ", tsd.TemperatureBackupValue)
	fmt.Println("TemperatureBackupUnit: ", tsd.TemperatureBackupUnit)
	fmt.Println("AltitudeValue: ", tsd.AltitudeValue)
	fmt.Println("AltitudeUnit: ", tsd.AltitudeUnit)
	fmt.Println("IlluminanceValue: ", tsd.IlluminanceValue)
	fmt.Println("IlluminanceUnit: ", tsd.IlluminanceUnit)
	fmt.Println("Timestamp: ", tsd.Timestamp)
}

var getDataCmd = &cobra.Command{
	Use:   "get_data",
	Short: "Poll data from the gRPC server",
	Long:  `Connect to the gRPC server and poll the time series data. Command used to test out that the server is running.`,
	Run: func(cmd *cobra.Command, args []string) {
		doGetData()
	},
}
