package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

type ApiServer struct {
	listenAddr string
}

type apiFunc func(w http.ResponseWriter, r *http.Request) error

type ApiError struct {
	Error string
}

func makeHttpHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

func NewApiServer(listenAddr string) *ApiServer {
	return &ApiServer{
		listenAddr: listenAddr,
	}
}

func (s *ApiServer) run() {
	router := mux.NewRouter()

	router.HandleFunc("/account", makeHttpHandleFunc(s.handleAccount))
}

func (s *ApiServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *ApiServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *ApiServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *ApiServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *ApiServer) handleTransfer(w http.ResponseWriter, r *http.Request) error {
	return nil
}
