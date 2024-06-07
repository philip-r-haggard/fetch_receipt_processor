package main

import (
	"encoding/json"
	"net/http"
	"strings"
	"sync"

	"github.com/google/uuid"
)

var (
	receipts = make(map[string]Receipt)
	points   = make(map[string]int)
	mu       sync.Mutex
)

func processReceipt(w http.ResponseWriter, r *http.Request) {
	var receipt Receipt
	if err := json.NewDecoder(r.Body).Decode(&receipt); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id := uuid.New().String()
	mu.Lock()
	points[id] = calculatePoints(receipt)
	receipts[id] = receipt
	mu.Unlock()

	json.NewEncoder(w).Encode(ProcessResponse{ID: id})
}

func getPoints(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/receipts/")
	id = strings.TrimSuffix(id, "/points")

	mu.Lock()
	points, ok := points[id]
	mu.Unlock()

	if !ok {
		http.Error(w, "Receipt not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(PointsResponse{Points: points})
}
