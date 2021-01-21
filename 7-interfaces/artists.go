package interfaces

import (
	"fmt"
	"time"
)

type Track struct {
	Title  string
	Name   string
	Year   int
	Album  string
	Length time.Duration
}

func length(d string) time.Duration {
	duration, err := time.ParseDuration(d)
	if err != nil {
		fmt.Println("error parsing duration", err)
		panic(d)
	}
	return duration
}

var tracks = []*Track{
	{"Wish you were here", "Pink Floyd", 1990, "Another brick in the wall", length("3m40s")},
	{"Mr. Sandman", "SYML", 2019, "Covers", length("3m24s")},
	{"Aerials", "System of Down", 2000, "System", length("4m5s")},
	{"Mr. Sandman", "SYML", 2012, "Live", length("3m24s")},
}

type byArtist []*Track

func (a byArtist) Len() int           { return len(a) }
func (a byArtist) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byArtist) Less(i, j int) bool { return a[i].Name < a[j].Name }

type byYear []*Track

func (a byYear) Len() int           { return len(a) }
func (a byYear) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byYear) Less(i, j int) bool { return a[i].Year < a[j].Year }
