package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

type Task struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Done      bool      `json:"done"`
	CreatedAt time.Time `json:"created_at"`
}

type TaskStore struct {
	mu    sync.RWMutex
	tasks map[string]*Task
	seq   int
}

func NewTaskStore() *TaskStore {
	return &TaskStore{tasks: make(map[string]*Task)}
}

func (s *TaskStore) Add(title string) *Task {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.seq++
	t := &Task{
		ID:        fmt.Sprintf("task-%d", s.seq),
		Title:     title,
		CreatedAt: time.Now(),
	}
	s.tasks[t.ID] = t
	return t
}

func (s *TaskStore) List() []*Task {
	s.mu.RLock()
	defer s.mu.RUnlock()
	out := make([]*Task, 0, len(s.tasks))
	for _, t := range s.tasks {
		out = append(out, t)
	}
	return out
}

func (s *TaskStore) Toggle(id string) (*Task, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	t, ok := s.tasks[id]
	if !ok {
		return nil, false
	}
	t.Done = !t.Done
	return t, true
}

func main() {
	store := NewTaskStore()

	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			json.NewEncoder(w).Encode(store.List())
		case http.MethodPost:
			var body struct{ Title string }
			json.NewDecoder(r.Body).Decode(&body)
			t := store.Add(body.Title)
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(t)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/tasks/toggle", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		t, ok := store.Toggle(id)
		if !ok {
			http.Error(w, "not found", http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(t)
	})

	log.Println("Task API running on :8090")
	log.Fatal(http.ListenAndServe(":8090", nil))
}
