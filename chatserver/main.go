package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

type Message struct {
	ID      int    `json:"id"`
	Sender  string `json:"sender"`
	Content string `json:"content"`
}

type ChatServer struct {
	mu       sync.Mutex
	messages []Message
}

func NewChatServer() *ChatServer {
	return &ChatServer{messages: make([]Message, 0)}
}

func (s *ChatServer) postMessage(w http.ResponseWriter, r *http.Request) {
	var msg Message
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	s.mu.Lock()
	msg.ID = len(s.messages) + 1
	s.messages = append(s.messages, msg)
	s.mu.Unlock()

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(msg)
}

func (s *ChatServer) getMessages(w http.ResponseWriter, r *http.Request) {
	s.mu.Lock()
	defer s.mu.Unlock()
	json.NewEncoder(w).Encode(s.messages)
}

func main() {
	srv := NewChatServer()

	http.HandleFunc("/api/message", srv.postMessage)
	http.HandleFunc("/api/messages", srv.getMessages)

	log.Println("Chat server running on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
