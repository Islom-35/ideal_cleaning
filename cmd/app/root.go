package app

import (
	"example.com/m/cmd/app/server"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var root = &cobra.Command{Use: "run-grpc"}

var userCmd = &cobra.Command{
	Use: "user-server",
	Run: func(cmd *cobra.Command, args []string) {
		server.RunGrpcUserServer()
	},
}

func Execute() {
	root.AddCommand(userCmd)

	if err := root.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error while executing CLI %s", err)
		os.Exit(1)
	}
}
