package cmd

import (
	"fmt"
	"net"
	"shopify-backend-challenge/v1/src/server"

	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use: "server",
	Short: "Starts CRUD service",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Listening on:", net.JoinHostPort("0.0.0.0", "8080"))
		server := server.NewServer(net.JoinHostPort("0.0.0.0", "8080"))
		return server.Run()
	},
}

func init() {
	RootCmd.AddCommand(serverCmd)
}
