// Source represents a video source

package panoptic

import (
//	"github.com/ziutek/glib"
	"github.com/revmischa/gst"
	"fmt"
	"os"
)

type CameraPipeline *gst.Pipeline
type StreamNewSource struct {
	Pad *gst.Pad
}

type StreamSource struct {
	SourceSelect chan *StreamNewSource
	URISrc *gst.Element
}

type RTPRelay struct {
	Source *StreamSource
	// Destination chan *RTPPacket
}

func checkElem(e *gst.Element, name string) {
	if e == nil {
		fmt.Fprintln(os.Stderr, "can't make element: ", name)
		os.Exit(1)
	}
}

// constructor
func NewStreamSource(uri string) *StreamSource {
	// create instance
	ss := &StreamSource{make(chan *StreamNewSource), nil}

	// create uridecoder to fetch and parse stream into elements
	uriDec := gst.ElementFactoryMake("uridecodebin", "uri-decoder")
	checkElem(uriDec, "uri-decoder")
	uriDec.SetProperty("uri", uri)
	//uriDec.SetProperty("expose-all-streams", false)  // we don't want everything, just video we can decode
	// callback when we get a stream from the uri decoder
	uriDec.ConnectNoi("pad-added", uriPadAdded, ss)
	
	ss.URISrc = uriDec

	return ss
}

// create stream from filename
func NewFileSource(filename string) *StreamSource {
	uri, err := gst.FilenameToURI(filename)
	if err != nil {
		// print err.message
		fmt.Printf("Error parsing filename " + filename)
		return nil
	}
	return NewStreamSource(uri)
}

func uriPadAdded(ss *StreamSource, uriNewPad *gst.Pad) {
	caps := uriNewPad.GetCurrentCaps()
	fmt.Println("New pad: ", uriNewPad.GetName())
	fmt.Println("  Caps: ", caps.String())
	caps.Unref()
	ss.SourceSelect <- &StreamNewSource{uriNewPad}
}

func (ss *StreamSource) MP4Decoder() *gst.Element {
	dec := gst.ElementFactoryMake("", "mp4v-decoder")
	return dec
}

func (ss *StreamSource) NewRTPRelay(destinationURI string) {
	rtpEncoder := gst.ElementFactoryMake("rtppay", "RTP relay repayloader") // ?
	checkElem(rtpEncoder, "RTP Relay")
	// relay := &RTPRelay{}
}
	
