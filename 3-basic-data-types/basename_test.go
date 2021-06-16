package __basic_data_types

import "testing"

func TestBaseName1(t *testing.T) {
	t.Run("example 1", func(t *testing.T) {
		got := basename1("abc/def/jkl.go")
		expected := "jkl"

		if got != expected {
			t.Errorf("got != expected; %s != %s", got, expected)
		}
	})

	t.Run("example 2", func(t *testing.T) {
		got := basename1("abc/def/jkl.abc.go")
		expected := "jkl.abc"

		if got != expected {
			t.Errorf("got != expected; %s != %s", got, expected)
		}
	})

	t.Run("example 3", func(t *testing.T) {
		got := basename1("abc/def/jkl")
		expected := "jkl"

		if got != expected {
			t.Errorf("got != expected; %s != %s", got, expected)
		}
	})
}

func TestBaseName2(t *testing.T) {
	t.Run("example 1", func(t *testing.T) {
		got := basename2("abc/def/jkl.go")
		expected := "jkl"

		if got != expected {
			t.Errorf("got != expected; %s != %s", got, expected)
		}
	})

	t.Run("example 2", func(t *testing.T) {
		got := basename2("abc/def/jkl.abc.go")
		expected := "jkl.abc"

		if got != expected {
			t.Errorf("got != expected; %s != %s", got, expected)
		}
	})

	t.Run("example 3", func(t *testing.T) {
		got := basename2("abc/def/jkl")
		expected := "jkl"

		if got != expected {
			t.Errorf("got != expected; %s != %s", got, expected)
		}
	})
}