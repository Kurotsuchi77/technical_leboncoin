package http

import (
	"encoding/json"
	"errors"
	"github.com/Kurotsuchi77/technical_leboncoin/fizzbuzz"
	"github.com/gorilla/mux"
	"net/http"
)

// Handler - Endpoint handler containing http Router and Service object
type Handler struct {
	Router  *mux.Router
	Service *fizzbuzz.Service
}

// NewHandler - Returns a pointer to a new endpoint Handler
func NewHandler(service *fizzbuzz.Service) *Handler {
	return &Handler{
		Service: service,
	}
}

// SetupRoutes - Creates http routes and bind them to Service's methods
func (handler *Handler) SetupRoutes() {
	handler.Router = mux.NewRouter()

	handler.Router.HandleFunc("/api/fizzbuzz/request", handler.GetFizzBuzz).
		Queries(
			"int1", "{int1:[0-9]+}",
			"int2", "{int2:[0-9]+}",
			"limit", "{limit:[0-9]+}",
			"str1", "{str1:.*}",
			"str2", "{str2:.*}",
		).
		Methods("GET")

	handler.Router.HandleFunc("/api/fizzbuzz/statistics", handler.GetMostRequested).
		Methods("GET")
}

// GetFizzBuzz - Writes fizzbuzz result
func (handler *Handler) GetFizzBuzz(w http.ResponseWriter, r *http.Request) {
	var req *fizzbuzz.Request
	var result *fizzbuzz.Result
	var err error

	vars := mux.Vars(r)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// Create Request object from GET parameters
	if req, err = handler.Service.CreateRequest(vars); err != nil {
		sendErrorMessage(w, "Cannot create request from parameters", err)
		return
	}
	// Perform actual fizzbuzz
	if result, err = handler.Service.GetFizzBuzz(req); err != nil {
		sendErrorMessage(w, "Cannot get fizzbuzz result", err)
		return
	}

	if err = json.NewEncoder(w).Encode(*result); err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
}

type statisticsResponse struct {
	Request fizzbuzz.Request
	Hits    int64
}

// GetMostRequested - Writes statistics about most requested fizzbuzz
func (handler *Handler) GetMostRequested(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if req, hit := handler.Service.GetMostRequested(); req != nil {
		if err := json.NewEncoder(w).Encode(statisticsResponse{Request: *req, Hits: hit}); err != nil {
			panic(err)
		}
		w.WriteHeader(http.StatusOK)
	} else {
		sendErrorMessage(w, "No statistics available", errors.New(""))
	}
}

type errorResponse struct {
	Message string
	Error   string
}

func sendErrorMessage(w http.ResponseWriter, message string, error error) {

	w.WriteHeader(http.StatusInternalServerError)
	if err := json.NewEncoder(w).Encode(errorResponse{Message: message, Error: error.Error()}); err != nil {
		panic(err)
	}
}
