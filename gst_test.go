package panoptic

import (
	"testing"
	"fmt"
)

/*
func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
*/

func TestGSTVersion(t *testing.T) {
	major, minor, micro, nano := GSTVersion()

	if (major != 1) {
		t.Error("Major version should be 1, got: " + string(major))
	}

	fmt.Printf("gstreamer version reported: %v.%v.%v.%v\n", major, minor, micro, nano)
}

func TestOK(t *testing.T) {
	print("OK\n")
}
