package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

type ApiFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string
}

func makeHTTPHandleFunc(f ApiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

type APIServer struct {
	listenAddr string
}

func CreateAPISever(listenAddr string) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
	}
}

func (s *APIServer) Run() {
	fmt.Println("running")
	log.Println("running")
	router := mux.NewRouter()
	router.HandleFunc("/patient", makeHTTPHandleFunc(s.handlePatient))
	log.Println("API server running on port: ", s.listenAddr)
	http.ListenAndServe(s.listenAddr, router)
}

func (s *APIServer) handlePatient(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		s.handleGetPatient(w, r)
	}
	if r.Method == "POST" {
		s.handleCreatePatient(w, r)
	}
	return fmt.Errorf("method not allowed %s", r.Method)
}

func (s *APIServer) handleGetPatient(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleCreatePatient(w http.ResponseWriter, r *http.Request) error {
	return nil
}
