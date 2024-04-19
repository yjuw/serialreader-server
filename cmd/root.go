package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "serialreader-server",
	Short: "Serve time-series data",
	Long:  `Serve time-series data from a connected Arduino device with an attached 'Sparkfun Weather Shield' device over gRPC.`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
