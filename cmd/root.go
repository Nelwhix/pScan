package cmd

import (
	"os"

	"github.com/spf13/cobra"
)


var rootCmd = &cobra.Command{
	Version: "0.0.2",
	Use:   "pScan",
	Short: "Fast TCP port scanner",
	Long: "pScan - short for Port Scanner - executes TCP port scan on a list of hosts.\npScan allows you to add, list, and delete hosts from the list.\npScan executes a port scan on specified TCP ports. You can customize the target ports using a command line flag.",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringP("hosts-file", "f", "pScan.hosts", "pScan hosts file")

	versionTemplate := `{{printf "%s: %s - version %s\n" .Name .Short .Version}}`
	rootCmd.SetVersionTemplate(versionTemplate)
}



