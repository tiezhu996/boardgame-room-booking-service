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
	var mu sync.Mutex
	success := 0
	for _, id := range roomIDs {
		wg.Add(1)
		go func(roomID string) {
			defer wg.Done()
			if s.store.Reserve(roomID) {
				mu.Lock()
				success++
				mu.Unlock()
			}
		}(id)
	}
	wg.Wait()
	return success
}
