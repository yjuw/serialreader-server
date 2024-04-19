package cmd // github.com/bartmika/serialreader-server/cmd/serve.go

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"

	server "github.com/yjuw/serialreader-server/internal"
)

var (
	port              int
	arduinoDevicePath string
)

func init() {
	// The following are required.
	serveCmd.Flags().StringVarP(&arduinoDevicePath, "arduino_path", "f", "/dev/cu.usbmodem14201", "The location of the connected arduino device on your computer.")
	serveCmd.MarkFlagRequired("arduino_path")

	// The following are optional and will have defaults placed when missing.
	serveCmd.Flags().IntVarP(&port, "port", "p", 50051, "The port to run this server on")

	// Make this sub-command part of our application.
	rootCmd.AddCommand(serveCmd)
}

func doServe() {
	// Setup our server.
	server := server.New(arduinoDevicePath, port)

	// DEVELOPERS CODE:
	// The following code will create an anonymous goroutine which will have a
	// blocking chan `sigs`. This blocking chan will only unblock when the
	// golang app receives a termination command; therfore the anyomous
	// goroutine will run and terminate our running application.
	//
	// Special Thanks:
	// (1) https://gobyexample.com/signals
	// (2) https://guzalexander.com/2017/05/31/gracefully-exit-server-in-go.html
	//
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs // Block execution until signal from terminal gets triggered here.
		server.StopMainRuntimeLoop()
	}()
	server.RunMainRuntimeLoop()
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run the gRPC server",
	Long:  `Run the gRPC server to allow other services to access the serial reader`,
	Run: func(cmd *cobra.Command, args []string) {
		doServe()
	},
}
