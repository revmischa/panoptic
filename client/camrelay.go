package main

import (
	"github.com/revmischa/panoptic"
	"github.com/revmischa/gst"
	"fmt"
)

/*
  this connects to an IP camera and receives an event on a channel when
  a stream has been made available

hayward~/dev/panoptic-client (master)? $ go run camrelay.go
got src pad:  /GstPipeline:camrelay/GstURIDecodeBin:uri-decoder.GstGhostPad:src_0

*/
func main() {
	// camera stream source
	ss := panoptic.NewStreamSource("rtsp://10.0.2.70/mpeg4/media.amp")

	// output (test)
	sink := gst.ElementFactoryMake("autovideosink", "video-output")

	// pipeline
	pl := gst.NewPipeline("camrelay")
	pl.Add(ss.URISrc, sink)
	pl.SetState(gst.STATE_PLAYING)

	go uriDecodeStatus(ss)

	panoptic.RunMainLoop()
}

func uriDecodeStatus(ss *panoptic.StreamSource) {
	for {
		select {
		case newSrc := <-(ss.SourceSelect):
			fmt.Println("got src pad: ", newSrc.Pad.GetPathString())
		}
	}
}
