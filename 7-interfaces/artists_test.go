package interfaces

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSortTracksByName(t *testing.T) {
	var tracks = []*Track{
		{"Mr. Sandman", "SYML", 2019, "Covers", length("3m24s")},
		{"Aerials", "System of Down", 2000, "System", length("4m5s")},
		{"Mr. Sandman", "SYML", 2012, "Live", length("3m24s")},
		{"Wish you were here", "Pink Floyd", 1990, "Another brick in the wall", length("3m40s")},
	}

	sort.Sort(byArtist(tracks))
	require.Equal(t, tracks[0].Name, "Pink Floyd")
	require.Equal(t, tracks[3].Name, "System of Down")
}

func TestSortTracksByYear(t *testing.T) {
	var tracks = []*Track{
		{"Mr. Sandman", "SYML", 2019, "Covers", length("3m24s")},
		{"Aerials", "System of Down", 2000, "System", length("4m5s")},
		{"Mr. Sandman", "SYML", 1989, "Live", length("3m24s")},
		{"Wish you were here", "Pink Floyd", 1990, "Another brick in the wall", length("3m40s")},
	}

	sort.Sort(byYear(tracks))
	require.Equal(t, tracks[0].Year, 1989)
	require.Equal(t, tracks[3].Year, 2019)
}
