package service

import (
	"errors"
	"github.com/TolkinSL/wb-l2/l2_18/internal/models"
	"sync"
	"time"
)

type CalendarService struct {
	mu     sync.RWMutex
	events map[int]models.Event
	nextID int
}

func NewCalendarService() *CalendarService {
	return &CalendarService{
		events: make(map[int]models.Event),
		nextID: 1,
	}
}

func (s *CalendarService) CreateEvent(e models.Event) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	e.ID = s.nextID
	s.events[e.ID] = e
	s.nextID++
	return e.ID, nil
}

func (s *CalendarService) UpdateEvent(e models.Event) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.events[e.ID]; !ok {
		return errors.New("event not found")
	}
	s.events[e.ID] = e
	return nil
}

func (s *CalendarService) DeleteEvent(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.events[id]; !ok {
		return errors.New("event not found")
	}
	delete(s.events, id)
	return nil
}

func (s *CalendarService) GetEventsForPeriod(userID int, start time.Time, days int) []models.Event {
	s.mu.RLock()
	defer s.mu.RUnlock()
	end := start.AddDate(0, 0, days)
	var result []models.Event
	for _, e := range s.events {
		if e.UserID == userID && !e.Date.Before(start) && e.Date.Before(end) {
			result = append(result, e)
		}
	}
	return result
}
