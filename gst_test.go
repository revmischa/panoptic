package panoptic

import (
	"testing"
	"github.com/ziutek/gst"
)

/*
func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
*/

func TestGSTVersion(t *testing.T) {
	major, minor, micro, nano := gst.GetVersion()

	if (major != 1) {
		t.Error("Major version should be 1, got: " + string(major))
	}
}

func TestOK(t *testing.T) {
	print("OK\n")
}
