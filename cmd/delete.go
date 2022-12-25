package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
	"github.com/Nelwhix/pScan/scan"
)

var deleteCmd = &cobra.Command{
	Use:   "delete <host1>...<host n>",
	Aliases: []string{"d"},
	Short: "Delete host(s) from list",
	SilenceUsage: true,
	Args: cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		hostsFile, err := cmd.Flags().GetString("hosts-file")
		if err != nil {
			return err 
		}
		return deleteAction(os.Stdout, hostsFile, args)
	},
}

func init() {
	hostsCmd.AddCommand(deleteCmd)
}

func deleteAction(out io.Writer, hostsFile string, args []string) error {
	hl := &scan.HostsList{}

	if err := hl.Load(hostsFile); err != nil {
		return err 
	}

	for _, h := range args {
		if err := hl.Remove(h); err != nil {
			return err 
		}

		fmt.Fprintln(out, "Deleted host:", h)
	}

	return hl.Save(hostsFile)
}