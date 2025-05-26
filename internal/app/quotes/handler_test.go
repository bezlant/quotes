package quotes

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

func setupRouter() *mux.Router {
	s := NewService()
	h := NewHandler(s)
	r := mux.NewRouter()
	h.RegisterRoutes(r)
	return r
}

func TestCreateQuoteHandler(t *testing.T) {
	t.Log("Setting up router")
	r := setupRouter()

	t.Log("Creating POST request with quote payload")
	body := `{"author":"Test","quote":"This is a test quote"}`
	req := httptest.NewRequest(http.MethodPost, "/quotes", strings.NewReader(body))
	w := httptest.NewRecorder()

	t.Log("Sending request to handler")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Fatalf("Expected 201 Created, got %d", w.Code)
	}

	var q Quote
	err := json.NewDecoder(w.Body).Decode(&q)
	if err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}
	t.Logf("Response quote: %+v", q)

	if q.Author != "Test" || q.Quote != "This is a test quote" {
		t.Errorf("Unexpected quote data: %+v", q)
	}
}

func TestGetQuotesHandler(t *testing.T) {
	t.Log("Setting up router")
	r := setupRouter()

	t.Log("Adding a quote via POST")
	reqBody := `{"author":"Alice","quote":"Hello"}`
	r.ServeHTTP(httptest.NewRecorder(), httptest.NewRequest("POST", "/quotes", strings.NewReader(reqBody)))

	t.Log("Sending GET request to /quotes")
	req := httptest.NewRequest("GET", "/quotes", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected 200 OK, got %d", w.Code)
	}

	var quotes []Quote
	err := json.NewDecoder(w.Body).Decode(&quotes)
	if err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}
	t.Logf("Fetched quotes: %+v", quotes)

	if len(quotes) != 1 {
		t.Fatalf("Expected 1 quote, got %d", len(quotes))
	}
}

func TestDeleteQuoteHandler(t *testing.T) {
	t.Log("Setting up router")
	r := setupRouter()

	t.Log("Creating a quote")
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(map[string]string{
		"author": "X",
		"quote":  "Y",
	})
	w := httptest.NewRecorder()
	r.ServeHTTP(w, httptest.NewRequest("POST", "/quotes", &buf))

	var created Quote
	json.NewDecoder(w.Body).Decode(&created)
	t.Logf("Created quote with ID: %d", created.ID)

	t.Log("Sending DELETE request")
	delReq := httptest.NewRequest("DELETE", "/quotes/"+strconv.Itoa(created.ID), nil)
	delResp := httptest.NewRecorder()
	r.ServeHTTP(delResp, delReq)

	if delResp.Code != http.StatusNoContent {
		t.Fatalf("Expected 204 No Content, got %d", delResp.Code)
	} else {
		t.Log("Quote deleted successfully")
	}
}
