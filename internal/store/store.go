package store

import (
	"fmt"
	"sync"

	"boardgame/internal/model"
)

type Store struct {
	mu    sync.Mutex
	rooms map[string]*model.Room
	seq   int
}

func New() *Store { return &Store{rooms: map[string]*model.Room{}} }

func (s *Store) AddRoom(r model.Room) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if r.ID == "" {
		return fmt.Errorf("room id required")
	}
	c := r
	s.rooms[r.ID] = &c
	return nil
}

// Reserve attempts to book one seat in a room; returns false when full.
func (s *Store) Reserve(roomID string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	r, ok := s.rooms[roomID]
	if !ok {
		return false
	}
	if !model.CanBook(r.Capacity, r.Booked) {
		return false
	}
	r.Booked++
	s.seq++
	return true
}

func (s *Store) BookedCount(roomID string) int {
	s.mu.Lock()
	defer s.mu.Unlock()
	r, ok := s.rooms[roomID]
	if !ok {
		return 0
	}
	return r.Booked
}

func (s *Store) TotalBooked() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	total := 0
	for _, r := range s.rooms {
		total += r.Booked
	}
	return total
}

func (s *Store) RoomCount() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return len(s.rooms)
}
