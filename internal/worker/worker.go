package worker

import (
	"boardgame/internal/store"
)

type Worker struct{ store *store.Store }

func New(s *store.Store) *Worker { return &Worker{store: s} }

func (w *Worker) Run() int {
	return w.store.TotalBooked() + 1
}
