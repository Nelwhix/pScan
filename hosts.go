package cmd

import (
	"github.com/spf13/cobra"
)


var hostsCmd = &cobra.Command{
	Use:   "hosts",
	Short: "Manage the hosts list",
	Long: `Manages the hosts lists for pScan
	
Add hosts with the add command
Delete hosts with the delete command
List hosts with the list command.`,
}

func init() {
	rootCmd.AddCommand(hostsCmd)
}
