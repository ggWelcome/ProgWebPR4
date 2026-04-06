package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"sync"
)

// Структури
type Generator struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Power  int    `json:"power"`
	Status string `json:"status"`
}

type Consumer struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Load   int    `json:"load"`
	Status string `json:"status"`
}

type Sensor struct {
	ID    int    `json:"id"`
	Type  string `json:"type"`
	Value int    `json:"value"`
	Unit  string `json:"unit"`
}

var (
	generators = make(map[int]Generator)
	consumers  = make(map[int]Consumer)
	sensors    = make(map[int]Sensor)

	genID, consID, sensID = 1, 1, 1
	mu                    sync.Mutex
)

// Універсальні хендлери
func handleGenerators(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		list := []Generator{}
		for _, g := range generators {
			list = append(list, g)
		}
		json.NewEncoder(w).Encode(list)
	case http.MethodPost:
		var g Generator
		if err := json.NewDecoder(r.Body).Decode(&g); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}
		mu.Lock()
		g.ID = genID
		genID++
		generators[g.ID] = g
		mu.Unlock()
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(g)
	case http.MethodDelete:
		idStr := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}
		mu.Lock()
		defer mu.Unlock()
		if _, ok := generators[id]; !ok {
			http.Error(w, "Not found", http.StatusNotFound)
			return
		}
		delete(generators, id)
		w.WriteHeader(http.StatusNoContent)
	}
}

func handleConsumers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		list := []Consumer{}
		for _, c := range consumers {
			list = append(list, c)
		}
		json.NewEncoder(w).Encode(list)
	case http.MethodPost:
		var c Consumer
		if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}
		mu.Lock()
		c.ID = consID
		consID++
		consumers[c.ID] = c
		mu.Unlock()
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(c)
	case http.MethodDelete:
		idStr := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}
		mu.Lock()
		defer mu.Unlock()
		if _, ok := consumers[id]; !ok {
			http.Error(w, "Not found", http.StatusNotFound)
			return
		}
		delete(consumers, id)
		w.WriteHeader(http.StatusNoContent)
	}
}

func handleSensors(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		list := []Sensor{}
		for _, s := range sensors {
			list = append(list, s)
		}
		json.NewEncoder(w).Encode(list)
	case http.MethodPost:
		var s Sensor
		if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}
		mu.Lock()
		s.ID = sensID
		sensID++
		sensors[s.ID] = s
		mu.Unlock()
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(s)
	case http.MethodDelete:
		idStr := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}
		mu.Lock()
		defer mu.Unlock()
		if _, ok := sensors[id]; !ok {
			http.Error(w, "Not found", http.StatusNotFound)
			return
		}
		delete(sensors, id)
		w.WriteHeader(http.StatusNoContent)
	}
}

func main() {
	http.HandleFunc("/generators", handleGenerators)
	http.HandleFunc("/consumers", handleConsumers)
	http.HandleFunc("/sensors", handleSensors)

	// Swagger UI
	http.Handle("/", http.FileServer(http.Dir(".")))

	http.ListenAndServe(":8080", nil)
}
