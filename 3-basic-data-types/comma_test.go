package __basic_data_types

import "testing"

func TestComma(t *testing.T) {
//	TODO: check out all methods offered by strings, bytes, strconv and unicode
	t.Run("example 1", func(t *testing.T) {
		num := "123456"
		expected := "123,456"

		got := comma(num)

		if got != expected {
			t.Errorf("got != expected; %s != %s", got, expected)
		}
	})

	t.Run("example 2", func(t *testing.T) {
		num := "123"
		expected := "123"

		got := comma(num)

		if got != expected {
			t.Errorf("got != expected; %s != %s", got, expected)
		}
	})

	t.Run("example 3", func(t *testing.T) {
		num := "12345"
		expected := "12,345"

		got := comma(num)

		if got != expected {
			t.Errorf("got != expected; %s != %s", got, expected)
		}
	})

	t.Run("example 4", func(t *testing.T) {
		num := "12345678900"
		expected := "12,345,678,900"

		got := comma(num)

		if got != expected {
			t.Errorf("got != expected; %s != %s", got, expected)
		}
	})
}