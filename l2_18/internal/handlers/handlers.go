package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/TolkinSL/wb-l2/l2_18/internal/models"
	"github.com/TolkinSL/wb-l2/l2_18/internal/service"
)

type Handler struct {
	Service *service.CalendarService
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("[%s] %s %s %v", r.Method, r.URL.Path, r.RemoteAddr, time.Since(start))
	})
}

func sendJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func sendError(w http.ResponseWriter, status int, msg string) {
	sendJSON(w, status, map[string]string{"error": msg})
}

func (h *Handler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		sendError(w, http.StatusMethodNotAllowed, "only POST allowed")
		return
	}

	userID, err := strconv.Atoi(r.FormValue("user_id"))
	date, errDate := time.Parse("2006-01-02", r.FormValue("date"))
	title := r.FormValue("title")

	if err != nil || errDate != nil || title == "" {
		sendError(w, 400, "invalid input data")
		return
	}

	id, _ := h.Service.CreateEvent(models.Event{
		UserID: userID,
		Date:   date,
		Title:  title,
	})
	sendJSON(w, 200, map[string]string{"result": "created, id: " + strconv.Itoa(id)})
}

func (h *Handler) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		sendError(w, http.StatusMethodNotAllowed, "only POST allowed")
		return
	}

	id, _ := strconv.Atoi(r.FormValue("id"))
	userID, _ := strconv.Atoi(r.FormValue("user_id"))
	date, errDate := time.Parse("2006-01-02", r.FormValue("date"))

	err := h.Service.UpdateEvent(models.Event{
		ID:     id,
		UserID: userID,
		Date:   date,
		Title:  r.FormValue("title"),
	})

	if errDate != nil {
		sendError(w, 400, "invalid date")
		return
	}

	if err != nil {
		sendError(w, 503, err.Error())
		return
	}
	sendJSON(w, 200, map[string]string{"result": "updated"})
}

func (h *Handler) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		sendError(w, http.StatusMethodNotAllowed, "only POST allowed")
		return
	}

	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		sendError(w, 400, "invalid id")
		return
	}

	if err := h.Service.DeleteEvent(id); err != nil {
		sendError(w, 503, err.Error())
		return
	}
	sendJSON(w, 200, map[string]string{"result": "deleted"})
}

func (h *Handler) getEvents(w http.ResponseWriter, r *http.Request, days int) {
	query := r.URL.Query()
	userID, _ := strconv.Atoi(query.Get("user_id"))
	date, err := time.Parse("2006-01-02", query.Get("date"))

	if err != nil {
		sendError(w, 400, "invalid date")
		return
	}

	events := h.Service.GetEventsForPeriod(userID, date, days)
	sendJSON(w, 200, map[string]interface{}{"result": events})
}

func (h *Handler) EventsForDay(w http.ResponseWriter, r *http.Request)   { h.getEvents(w, r, 1) }
func (h *Handler) EventsForWeek(w http.ResponseWriter, r *http.Request)  { h.getEvents(w, r, 7) }
func (h *Handler) EventsForMonth(w http.ResponseWriter, r *http.Request) { h.getEvents(w, r, 30) }
