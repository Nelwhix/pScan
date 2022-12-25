package cmd

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"strconv"
	"os"
	"strings"
	"testing"

	"github.com/Nelwhix/pScan/scan"
)

func setup(t *testing.T, hosts []string, initList bool) (string, func()) {
	tf, err := os.CreateTemp("", "pScan")
	if err != nil {
		t.Fatal(err)
	}

	

	if initList {
		hl := &scan.HostsList{}

		for _, h := range hosts {
			hl.Add(h)
		}

		if err := hl.Save(tf.Name()); err != nil {
			t.Fatal(err)
		}
	}
	tf.Close()
	return tf.Name(), func() {
		os.Remove(tf.Name())
	}
}

func TestHostActions(t *testing.T) {
	hosts := []string{
		"host1",
		"host2",
		"host3",
	}

	testCases := []struct{
		name string
		args []string
		expectedOutput string
		initList bool
		actionFunction func (io.Writer, string, []string) error
	} {
		{
			name: "AddAction",
			args: hosts,
			expectedOutput: "Added host: host1\nAdded host: host2\nAdded host: host3\n",
			initList: false,
			actionFunction: addAction,
		},
		{
			name: "ListAction",
			expectedOutput: "host1\nhost2\nhost3\n",
			initList: true,
			actionFunction: listAction,
		},
		{
			name: "DeleteAction",
			args: []string{"host1", "host2"},
			expectedOutput: "Deleted host: host1\nDeleted host: host2\n",
			initList: true,
			actionFunction: deleteAction,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tf, cleanup := setup(t, hosts, tc.initList)
			defer cleanup()

			var out bytes.Buffer

			if err := tc.actionFunction(&out, tf, tc.args); err != nil {
				t.Fatalf("Expected no error, got %q\n", err)
			}

			if out.String() != tc.expectedOutput {
				t.Errorf("Expected output %q, got %q\n", tc.expectedOutput, out.String())
			}
		})
	}
}

func TestIntegration(t *testing.T) {
	hosts := []string{
		"host1",
		"host2",
		"host3",
	}

	tf, cleanup := setup(t, hosts, false)
	defer cleanup()

	delHost := "host2"

	hostsEnd := []string{
		"host1",
		"host3",
	}

	var out bytes.Buffer
	expectedOut := ""

	for _, v := range hosts {
		expectedOut += fmt.Sprintf("Added host: %s\n", v)
	}

	expectedOut += strings.Join(hosts, "\n")
	expectedOut += fmt.Sprintln()
	expectedOut += fmt.Sprintf("Deleted host: %s\n", delHost)
	expectedOut += strings.Join(hostsEnd, "\n")
	expectedOut += fmt.Sprintln()

	if err := addAction(&out, tf, hosts); err != nil {
		t.Fatalf("Expected no error, got %q\n", err)
	}

	if err := listAction(&out, tf, nil); err != nil {
		t.Fatalf("Expected no error, got %q\n", err)
	}

	if err := deleteAction(&out, tf, []string{delHost}); err != nil {
		t.Fatalf("Expected no error, got %q\n", err)
	}

	if err := listAction(&out, tf, nil); err != nil {
		t.Fatalf("Expected no error, got %q\n", err)
	}

	if out.String() != expectedOut {
		t.Errorf("Expected output %q, got %q\n", expectedOut, out.String())
	}
}

func TestScanAction(t *testing.T) {
	hosts := []string{
		"127.0.0.1",
		"192.168.0.1",
	}

	tf, cleanup := setup(t, hosts, true)

	defer cleanup()
	ports := []int{}

	for i := 0; i < 2; i++ {
		ln, err := net.Listen("tcp", net.JoinHostPort("localhost", "0"))
		if err != nil {
			t.Fatal(err)
		}
		defer ln.Close()

		_, portStr, err := 
	}
}