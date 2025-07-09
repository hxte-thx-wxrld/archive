package main

import (
	"github.com/32leaves/bel"
	"github.com/htw-archive/pkg/model"
)

func main() {
	OutputStruct(model.PaginatedArtistLookup{})
	OutputStruct(model.Artist{})
	OutputStruct(model.Music{})
	OutputStruct(model.Release{})
	OutputStruct(model.MusicInRelease{})

	OutputStruct(model.InboxItem{})
	OutputStruct(model.PaginatedInboxItems{})
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
