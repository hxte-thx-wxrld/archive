package main

import (
	"github.com/32leaves/bel"
	"github.com/htw-archive/pkg/api"
)

func main() {
	OutputStruct(api.Artist{})
	OutputStruct(api.Release{})
	OutputStruct(api.MusicRow{})
	OutputStruct(api.MusicInRelease{})
	OutputStruct(api.TrackEditRequest{})
}

func OutputStruct(s interface{}) {
	ts, err := bel.Extract(s)
	if err != nil {
		panic(err)
	}

	err = bel.Render(ts)
	if err != nil {
		panic(err)
	}
}
