package cmd

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
	server "github.com/yjuw/serialreader-server/internal"
)

var (
	port int
)

func init() {
	serveCmd.Flags().IntVarP(&port, "port", "p", 50051, "The port to run this server on")
	rootCmd.AddCommand(serveCmd)
}

func doServe() {
	server := server.New(port)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs //Block execution until signal from terminal gets triggered
		server.StopMainRuntimeLoop()
	}()
	server.RunMainRuntimeLoop()
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run the gRPC server",
	Long:  `Run the gRPC server to allow other sevices to access the serial reader`,
	Run: func(cmd *cobra.Command, args []string) {
		doServe()
	},
}
