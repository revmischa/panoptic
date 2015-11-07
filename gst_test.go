package panoptic

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"github.com/revmischa/gst"
)

// utilities
func getVideoStream(t *testing.T) *StreamSource {
	return NewFileSource("./t/test.mp4")
}

// func (t *testing.T)getH264Video() gst.Element {
// 	ss := t.getVideoStream()
// 	srcElement = ss.URISrc

// 	// c

// 	// link up qtdemux and extract the h264 video stream
// 	demuxBin, err := ParseBinFromDescription("qtdemux name=demuxer  demuxer.video_0")
// 	srcElement.Link(demuxBin)


// }

func getH264RTPSender(t *testing.T) (*gst.Bin) {
	launch := "filesrc location=./t/test.mp4 ! qtdemux name=demuxer  demuxer.video_0 ! rtph264pay config-interval=1 ! udpsink host=127.0.0.1 port=5000"
	rtpSendBin, err := gst.ParseBinFromDescription(launch)
	if rtpSendBin == nil {
		t.Fatal("Error parsing bin (" + err.Error() + "): " + launch)
	}
	return rtpSendBin
}

func getH264RTPReceiver(t *testing.T) (*gst.Bin) {
	launch := "udpsrc uri=udp://127.0.0.1:5000 caps=\"application/x-rtp,media=(string)video\" ! rtph264depay ! decodebin ! autovideosink sync=false"
	rtpPlayBin, err := gst.ParseBinFromDescription(launch)
	if rtpPlayBin == nil {
		t.Fatal("Error parsing bin (" + err.Error() + "): " + launch)
	}
	return rtpPlayBin
}
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

func TestRTPRelay(t *testing.T) {
	// make a video source, send it via RTP to localhost, depayload it

	// source:
	// filesrc location=t/test.mp4 ! qtdemux name=demuxer  demuxer.video_0 ! rtph264pay config-interval=1 ! udpsink host=127.0.0.1 port=5000
	// sink:
	// udpsrc uri=udp://127.0.0.1:5000 caps="application/x-rtp,media=(string)video" ! rtpjitterbuffer ! rtph264depay ! decodebin ! autovideosink sync=false

	sender := getH264RTPSender(t)
	receiver := getH264RTPReceiver(t)
	assert.NotNil(t, sender, "Constructed RTP sender")
	assert.NotNil(t, receiver, "Constructed RTP receiver")

	senderPipeline := gst.NewPipeline("rtpSender")
	senderPipeline.Add(sender.AsElement())
	receiverPipeline := gst.NewPipeline("rtpReceiver")
	receiverPipeline.Add(receiver.AsElement())


	receiverPipeline.SetState(gst.STATE_PLAYING)
	// senderPipeline.SetState(gst.STATE_PLAYING)
	receiverPipeline.GetBus().AddSignalWatch()
	receiverPipeline.GetBus().Connect("message::error", gstError, nil)

	RunMainLoop()
}

func gstError(pipeline *gst.Pipeline, msg *gst.Message) {
	gerr, dbg := msg.ParseError()
	panic(fmt.Sprintf("Got gstreamer error: %s\n[%s]\n", gerr.Error(), dbg))
}