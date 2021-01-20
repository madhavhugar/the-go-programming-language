package interfaces

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIntegersSort(t *testing.T) {
	input := integers{7, 3, 9, 2, 3, 1, 2, 10, 8}

	sort.Sort(input)
	expected := integers{1, 2, 2, 3, 3, 7, 8, 9, 10}
	require.Equal(t, expected, input)
}

func TestIntegersReverseSort(t *testing.T) {
	input := integers{7, 3, 9, 2, 3, 1, 2, 10, 8}

	sort.Sort(sort.Reverse(input))
	expected := integers{10, 9, 8, 7, 3, 3, 2, 2, 1}
	require.Equal(t, expected, input)
}
