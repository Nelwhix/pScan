package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/Nelwhix/pScan/scan"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:          "add <host1>...<hostn>",
	Aliases:      []string{"a"},
	Args:         cobra.MinimumNArgs(1),
	SilenceUsage: true,
	Short:        "Add new host(s) to list",
	RunE: func(cmd *cobra.Command, args []string) error {
		hostsFile, err := cmd.Flags().GetString("hosts-file")
		if err != nil {
			return err
		}

		return addAction(os.Stdout, hostsFile, args)
	},
}

func addAction(out io.Writer, hostsFile string, args []string) error {
	hl := &scan.HostsList{}
	if err := hl.Load(hostsFile); err != nil {
		return err
	}

	for _, h := range args {
		if err := hl.Add(h); err != nil {
			return err
		}
		_, err := fmt.Fprintln(out, "Added host:", h)
		if err != nil {
			return err
		}
	}

	return hl.Save(hostsFile)
}
func init() {
	hostsCmd.AddCommand(addCmd)
}
