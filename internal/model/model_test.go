package model

import "testing"

func TestCanBook(t *testing.T) {
	cases := []struct {
		capacity, booked int
		want             bool
	}{
		{4, 0, true},
		{4, 3, true},
		{4, 4, false},
		{4, 5, false},
	}
	for _, c := range cases {
		if got := CanBook(c.capacity, c.booked); got != c.want {
			t.Errorf("CanBook(%d,%d)=%v want %v", c.capacity, c.booked, got, c.want)
		}
	}
}
