package panoptic

import (
	    "github.com/ziutek/gst"
)

func GSTVersion() (int, int, int, int) {
	return gst.GetVersion()
}
