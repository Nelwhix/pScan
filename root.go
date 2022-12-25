/*
Copyright © 2022 Nelson Isioma
Copyrights apply to this source code.
Check LICENSE for details.

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)



// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Version: "0.0.1",
	Use:   "pScan",
	Short: "Fast TCP port scanner",
	Long: "pScan - short for Port Scanner - executes TCP port scan on a list of hosts.\npScan allows you to add, list, and delete hosts from the list.\npScan executes a port scan on specified TCP ports. You can customize the target ports using a command line flag.",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
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


