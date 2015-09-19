package panoptic

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"github.com/revmischa/gst"
)

/*
func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
*/

func TestGSTVersion(t *testing.T) {
	assert := assert.New(t)
	major, minor, micro, nano := GSTVersion()

	assert.Equal(major, 1, "Major version should be 1, got: "+string(major))

	fmt.Printf("gstreamer version reported: %v.%v.%v.%v\n", major, minor, micro, nano)
}

func TestOK(t *testing.T) {
	print("OK\n")
}

func TestURI(t *testing.T) {
	assert := assert.New(t)

	// valid file path
	uri, err := gst.FilenameToURI("./t/test.mp4")
	assert.Contains(uri, "panoptic/t/test.mp4", "FilenameToURI handles relative path")
	assert.Nil(err, "No error converting filename to URI")
}