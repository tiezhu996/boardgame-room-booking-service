package worker

import (
	"testing"

	"boardgame/internal/model"
	"boardgame/internal/store"
)

func TestRun(t *testing.T) {
	s := store.New()
	_ = s.AddRoom(model.Room{ID: "r1", Capacity: 2})
	s.Reserve("r1")
	if got := New(s).Run(); got != 1 {
		t.Fatalf("total=%d want 1", got)
	}
}
