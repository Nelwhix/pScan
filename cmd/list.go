/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/Nelwhix/pScan/scan"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Aliases: []string{"l"},
	Use:     "list",
	Short:   "List hosts in hosts list",
	RunE: func(cmd *cobra.Command, args []string) error {
		hostsFile, err := cmd.Flags().GetString("hosts-file")

		if err != nil {
			return err
		}

		return listAction(os.Stdout, hostsFile, args)
	},
}

func init() {
	hostsCmd.AddCommand(listCmd)
}

func listAction(out io.Writer, hostsFile string, args []string) error {
	hl := &scan.HostsList{}

	if err := hl.Load(hostsFile); err != nil {
		return err
	}

	if len(hl.Hosts) == 0 {
		fmt.Fprintln(out, "no hosts added")
	}

	for _, h := range hl.Hosts {
		if _, err := fmt.Fprintln(out, h); err != nil {
			return err
		}
	}

	return nil
}
