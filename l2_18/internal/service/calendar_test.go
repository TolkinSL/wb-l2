package service

import (
	"testing"
	"time"
	"github.com/TolkinSL/wb-l2/l2_18/internal/models"
)

func TestCalendarWorkflow(t *testing.T) {
    s := NewCalendarService()
    date, _ := time.Parse("2006-01-02", "2026-05-20")

    id, _ := s.CreateEvent(models.Event{UserID: 1, Date: date, Title: "Test"})
    
    events := s.GetEventsForPeriod(1, date, 1)
    if len(events) != 1 || events[0].ID != id {
        t.Errorf("Событие не найдено")
    }

    err := s.DeleteEvent(id)
    if err != nil {
        t.Errorf("Ошибка при удалении: %v", err)
    }

    eventsAfter := s.GetEventsForPeriod(1, date, 1)
    if len(eventsAfter) != 0 {
        t.Errorf("Событие все еще существует после удаления")
    }
}