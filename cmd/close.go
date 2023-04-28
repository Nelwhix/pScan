package cmd

import (
	"io"
	"os"
	"fmt"

	"github.com/Nelwhix/pScan/scan"
	"github.com/spf13/cobra"
)

var closeCmd = &cobra.Command{
	Use:   "close",
	Short: "For closing open TCP ports",
	RunE: func(cmd *cobra.Command, args []string) error {
		ports, err := cmd.Flags().GetIntSlice("ports")

		if err != nil {
			return err 
		}

		return closeAction(os.Stdout, ports)
	},
}

func init() {
	rootCmd.AddCommand(closeCmd)

	closeCmd.Flags().IntSliceP("ports", "p", []int{}, "ports to scan")
}

func closeAction(out io.Writer, ports []int) error {
	message := ""
	hl := &scan.HostsList{[]string{"localhost"}}

	results := scan.Run(hl, ports)

	for _, r := range results {
		for _, p := range r.PortStates {
			if (!p.Open) {
				message += fmt.Sprintf("%s:%d is already closed\n", r.Host, p.Port)
			}
		}
	}

	_, err := fmt.Fprint(out, message)
	return err
}
