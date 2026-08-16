package store

import (
	"testing"

	"boardgame/internal/model"
)

func TestReserveRejectsWhenFull(t *testing.T) {
	s := New()
	_ = s.AddRoom(model.Room{ID: "r1", Capacity: 2})
	if !s.Reserve("r1") || !s.Reserve("r1") {
		t.Fatal("first two reservations should succeed")
	}
	if s.Reserve("r1") {
		t.Fatal("third reservation should be rejected when full")
	}
	if s.BookedCount("r1") != 2 {
		t.Fatalf("booked=%d want 2", s.BookedCount("r1"))
	}
}
