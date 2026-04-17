package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/TolkinSL/wb-l2/l2_18/internal/handlers"
	"github.com/TolkinSL/wb-l2/l2_18/internal/service"
)

func main() {
	port := flag.String("port", "8080", "HTTP port to listen on")
	flag.Parse()

	if envPort := os.Getenv("PORT"); envPort != "" {
		*port = envPort
	}

	svc := service.NewCalendarService()
	h := &handlers.Handler{Service: svc}

	mux := http.NewServeMux()

	mux.HandleFunc("/create_event", h.CreateEvent)
	mux.HandleFunc("/update_event", h.UpdateEvent)
	mux.HandleFunc("/delete_event", h.DeleteEvent)

	mux.HandleFunc("/events_for_day", h.EventsForDay)
	mux.HandleFunc("/events_for_week", h.EventsForWeek)
	mux.HandleFunc("/events_for_month", h.EventsForMonth)

	finalHandler := handlers.LoggingMiddleware(mux)

	log.Printf("Calendar server started on port %s", *port)
	if err := http.ListenAndServe(":"+*port, finalHandler); err != nil {
		log.Fatalf("Server failed: %v", err)
	}

}
