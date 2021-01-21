package interfaces

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSliceSort(t *testing.T) {
	ss := sort.StringSlice{"kitty", "madkitty", "meerkitty", "maddy"}

	expected := sort.StringSlice{"kitty", "maddy", "madkitty", "meerkitty"}
	sort.Sort(ss)

	require.Equal(t, expected, ss)
}

func TestSliceSortReverse(t *testing.T) {
	ss := sort.StringSlice{"kitty", "madkitty", "meerkitty", "maddy"}

	expected := sort.StringSlice{"meerkitty", "madkitty", "maddy", "kitty"}
	sort.Sort(sort.Reverse(ss))

	require.Equal(t, expected, ss)
}

func TestIntSort(t *testing.T) {
	input := sort.IntSlice{7, 3, 9, 2, 3, 1, 2, 10, 8}

	sort.Sort(input)
	expected := sort.IntSlice{1, 2, 2, 3, 3, 7, 8, 9, 10}

	require.Equal(t, expected, input)
}

func TestIntSortReverse(t *testing.T) {
	input := sort.IntSlice{7, 3, 9, 2, 3, 1, 2, 10, 8}

	sort.Sort(sort.Reverse(input))
	expected := sort.IntSlice{10, 9, 8, 7, 3, 3, 2, 2, 1}

	require.Equal(t, expected, input)
}
