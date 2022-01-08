package cmd

import (
	"fmt"
	"net"
	"shopify-backend-challenge/v1/src/csv"

	"github.com/spf13/cobra"
)

var csvCmd = &cobra.Command{
	Use: "csv",
	Short: "Starts CRUD service",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Listening on:", net.JoinHostPort("0.0.0.0", "8081"))
		server := csv.NewServer(net.JoinHostPort("0.0.0.0", "8081"))
		return server.Run()
	},
}

func init() {
	RootCmd.AddCommand(csvCmd)
}

