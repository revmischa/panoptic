package panoptic

import (
	"github.com/revmischa/gst"
	"github.com/ziutek/glib"
)

func GSTVersion() (int, int, int, int) {
	return gst.GetVersion()
}

func RunMainLoop() {
	glib.NewMainLoop(nil).Run()
}
