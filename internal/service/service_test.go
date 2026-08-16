package service

import (
	"testing"

	"boardgame/internal/model"
	"boardgame/internal/store"
)

func TestBatchReserve(t *testing.T) {
	st := store.New()
	_ = st.AddRoom(model.Room{ID: "r1", Capacity: 2})
	_ = st.AddRoom(model.Room{ID: "r2", Capacity: 2})
	s := New(st)
	ids := []string{"r1", "r1", "r2", "r2", "r1"}
	if got := s.BatchReserve(ids); got != 4 {
		t.Fatalf("reserved=%d want 4", got)
	}
	if st.TotalBooked() != 4 {
		t.Fatalf("total=%d want 4", st.TotalBooked())
	}
}
