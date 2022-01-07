package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// base command that executes
var RootCmd = &cobra.Command {
	Use: "root",
	Short: "This start everything",
	Long: "This start everything",
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		os.Exit(-1)		
	}
}

func init() {}
