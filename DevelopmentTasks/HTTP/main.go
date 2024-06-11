package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Event struct {
	UserID int       `json:"user_id"`
	Date   time.Time `json:"date"`
	Title  string    `json:"title"`
}

var events = make(map[int]Event)
var nextID = 1

func writeJSONResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func createEvent(w http.ResponseWriter, r *http.Request) {
	var event Event
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		writeJSONResponse(w, map[string]string{"error": "invalid request body"}, http.StatusBadRequest)
		return
	}

	eventID := nextID
	nextID++
	events[eventID] = event
	writeJSONResponse(w, map[string]interface{}{"result": "event created", "event_id": eventID}, http.StatusCreated)
}

func updateEvent(w http.ResponseWriter, r *http.Request) {
	var event Event
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		writeJSONResponse(w, map[string]string{"error": "invalid request body"}, http.StatusBadRequest)
		return
	}

	eventIDStr := r.URL.Query().Get("event_id")
	eventID, err := strconv.Atoi(eventIDStr)
	if err != nil || eventID < 1 {
		writeJSONResponse(w, map[string]string{"error": "invalid event_id"}, http.StatusBadRequest)
		return
	}

	if _, exists := events[eventID]; !exists {
		writeJSONResponse(w, map[string]string{"error": "event not found"}, http.StatusNotFound)
		return
	}

	events[eventID] = event
	writeJSONResponse(w, map[string]string{"result": "event updated"}, http.StatusOK)
}

func deleteEvent(w http.ResponseWriter, r *http.Request) {
	eventIDStr := r.URL.Query().Get("event_id")
	eventID, err := strconv.Atoi(eventIDStr)
	if err != nil || eventID < 1 {
		writeJSONResponse(w, map[string]string{"error": "invalid event_id"}, http.StatusBadRequest)
		return
	}

	delete(events, eventID)
	writeJSONResponse(w, map[string]string{"result": "event deleted"}, http.StatusOK)
}

func getEventsForDay(w http.ResponseWriter, r *http.Request) {
	dateStr := r.URL.Query().Get("date")
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		writeJSONResponse(w, map[string]string{"error": "invalid date"}, http.StatusBadRequest)
		return
	}

	var result []Event
	for _, event := range events {
		if event.Date.Year() == date.Year() && event.Date.YearDay() == date.YearDay() {
			result = append(result, event)
		}
	}

	writeJSONResponse(w, result, http.StatusOK)
}

func getEventsForWeek(w http.ResponseWriter, r *http.Request) {
	dateStr := r.URL.Query().Get("date")
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		writeJSONResponse(w, map[string]string{"error": "invalid date"}, http.StatusBadRequest)
		return
	}

	year, week := date.ISOWeek()
	var result []Event
	for _, event := range events {
		eventYear, eventWeek := event.Date.ISOWeek()
		if eventYear == year && eventWeek == week {
			result = append(result, event)
		}
	}

	writeJSONResponse(w, result, http.StatusOK)
}

func getEventsForMonth(w http.ResponseWriter, r *http.Request) {
	dateStr := r.URL.Query().Get("date")
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		writeJSONResponse(w, map[string]string{"error": "invalid date"}, http.StatusBadRequest)
		return
	}

	var result []Event
	for _, event := range events {
		if event.Date.Year() == date.Year() && event.Date.Month() == date.Month() {
			result = append(result, event)
		}
	}

	writeJSONResponse(w, result, http.StatusOK)
}

func main() {
	http.HandleFunc("/create_event", createEvent)
	http.HandleFunc("/update_event", updateEvent)
	http.HandleFunc("/delete_event", deleteEvent)
	http.HandleFunc("/events_for_day", getEventsForDay)
	http.HandleFunc("/events_for_week", getEventsForWeek)
	http.HandleFunc("/events_for_month", getEventsForMonth)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
