package cmd

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/Nelwhix/pScan/entity"
	"github.com/Nelwhix/pScan/utils"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

var process *entity.Process

var killCmd = &cobra.Command{
	Use:   "kill",
	Short: "kill the process on your local machine running on the specified port",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		port, err := cmd.Flags().GetInt("port")
		if err != nil {
			return err
		}

		process, err = findProcessRunningOnPort(port)
		if err != nil {
			return err
		}

		_, err = fmt.Fprintf(os.Stdout, "Are you sure you want to kill process: %v %v on port: %v? [Y/n]", process.Command, process.Name, port)
		if err != nil {
			return err
		}

		reader := bufio.NewReader(os.Stdin)
		for {
			s, _ := reader.ReadString('\n')
			s = strings.TrimSuffix(s, "\n")
			s = strings.ToLower(s)
			if len(s) > 1 {
				_, err = fmt.Fprintln(os.Stderr, "Please enter Y or N")
				if err != nil {
					return err
				}
				continue
			}
			if s == "n" {
				cmd.RunE = func(cmd *cobra.Command, args []string) error {
					_, err = fmt.Fprintln(os.Stdout, "Action aborted!")
					if err != nil {
						return err
					}

					return nil
				}

				return nil
			} else if s == "y" {
				break
			} else {
				continue
			}
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		err := killPort(process.PID)
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(killCmd)

	killCmd.Flags().IntP("port", "p", 0, "port to kill")
}

func findProcessRunningOnPort(port int) (*entity.Process, error) {
	portString := strconv.Itoa(port)
	cmd := exec.Command("lsof", "-i", ":"+portString)

	var buf bytes.Buffer
	cmd.Stdout = &buf
	err := cmd.Run()
	if err != nil {
		return nil, err
	}

	process := utils.ParseProcessOutput(buf.String())
	if err != nil {
		return nil, err
	}

	return &process, nil
}

func killPort(processId int) error {
	cmd := exec.Command("kill", strconv.Itoa(processId))

	var buf bytes.Buffer
	cmd.Stdout = &buf
	err := cmd.Run()
	if err != nil {
		return err
	}

	if buf.Len() == 0 {
		_, err = fmt.Fprintln(os.Stdout, fmt.Sprintf("killed process with id: %v", processId))

		if err != nil {
			return err
		}
	}

	return nil
}
