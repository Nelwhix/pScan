package scan_test

import (
	"errors"
	"os"
	"testing"

	"github.com/Nelwhix/pScan/scan"
)

func TestAdd(t *testing.T) {
	testCases := []struct {
		name string
		host string
		expectedLength int
		expectedErr error
	} {
		{"it can add new host", "host2", 2, nil},
		{"adding an existing host throws an error", "host1", 1, scan.ErrExists},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			hl := &scan.HostsList{}

			if err := hl.Add("host1"); err != nil {
				t.Fatal(err)
			}

			err := hl.Add(tc.host)

			if tc.expectedErr != nil {
				if err == nil {
					t.Fatalf("Expected error, got nil instead\n")
				}
				if !errors.Is(err, tc.expectedErr) {
					t.Errorf("Expected error %q, got %q instead\n", tc.expectedErr, err)
				}

				return
			}

			if err != nil {
				t.Fatalf("Expected no error, got %q instead\n", err)
			}

			if len(hl.Hosts) != tc.expectedLength {
				t.Errorf("Expected list length %d, got %d instead\n",
					tc.expectedLength, len(hl.Hosts))
			}

			if hl.Hosts[1] != tc.host {
				t.Errorf("Expected host name %q as index 1, got %q instead\n",
					tc.host, hl.Hosts[1])		
			}
		})
	}
}

func TestRemove(t *testing.T) {
	testCases := []struct {
		name string
		host string
		expectedLength int
		expectedError error 
	} {
		{"it removes a host", "host1", 1, nil},
		{"removing an unset host throws error", "host3", 1, scan.ErrNotExists},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			hl := &scan.HostsList{}

			for _, h := range []string{"host1", "host2"} {
				if err := hl.Add(h); err != nil {
					t.Fatal(err)
				}
			}

			err := hl.Remove(tc.host)

			if tc.expectedError != nil {
				if err == nil {
					t.Fatalf("Expected error, got nil instead\n")
				}

				if !errors.Is(err, tc.expectedError) {
					t.Errorf("Expected error %q, got %q instead\n",
						tc.expectedError, err)
				}

				return
			}

			if err != nil {
				t.Fatalf("Expected no error, got %q instead\n", err)
			}

			if len(hl.Hosts) != tc.expectedLength {
				t.Errorf("Expected list length %d, got %d instead\n", tc.expectedLength, len(hl.Hosts))
			}

			if hl.Hosts[0] == tc.host {
				t.Errorf("Host name %q should not be in the list\n", tc.host)
			}
		})
	}
}

func TestSaveLoad(t *testing.T) {
	hl1 := scan.HostsList{}
	hl2 := scan.HostsList{}

	hostName := "host1"
	hl1.Add(hostName)

	tf, err := os.CreateTemp("", "")

	if err != nil {
		t.Fatalf("Error creating temp file: %s", err)
	}
	defer os.Remove(tf.Name())

	if err := hl1.Save(tf.Name()); err != nil {
		t.Fatalf("Error saving list to file: %s", err)
	}

	if err := hl2.Load(tf.Name()); err != nil {
		t.Fatalf("Error getting list from file: %s", err)
	}

	if hl1.Hosts[0] != hl2.Hosts[0] {
		t.Errorf("Host %q should match %q host.", hl1.Hosts[0], hl2.Hosts[0])
	}
}

func TestLoadNoFile(t *testing.T) {
	tf, err := os.CreateTemp("", "")

	if err != nil {
		t.Fatalf("Error creating temp file: %s", err)
	}
	tf.Close()
	if err := os.Remove(tf.Name());  err != nil {
		t.Fatalf("Error deleting temp file: %s", err)
	}

	hl := &scan.HostsList{}

	if err := hl.Load(tf.Name()); err != nil {
		t.Errorf("Expected no error, got %q instead\n", err)
	}
}