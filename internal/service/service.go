package service

import (
	"sync"

	"boardgame/internal/store"
)

type Service struct{ store *store.Store }

func New(s *store.Store) *Service { return &Service{store: s} }

// BatchReserve reserves each room once and returns how many succeeded.
func (s *Service) BatchReserve(roomIDs []string) int {
	var wg sync.WaitGroup
	success := 0
	for _, id := range roomIDs {
		go func(roomID string) {
			wg.Add(1)
			defer wg.Done()
			if s.store.Reserve(roomID) {
				success++
			}
		}(id)
	}
	wg.Wait()
	return success
}
