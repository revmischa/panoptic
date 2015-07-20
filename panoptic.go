package panoptic

import (
	"github.com/revmischa/gst"
)

func GSTVersion() (int, int, int, int) {
	return gst.GetVersion()
}
