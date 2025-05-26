package quotes

import (
	"testing"
)

func TestAddQuote(t *testing.T) {
	s := NewService()
	q := s.AddQuote("Author1", "Quote1")

	if q.ID != 1 {
		t.Errorf("Expected ID to be 1, got %d", q.ID)
	}
	if q.Author != "Author1" || q.Quote != "Quote1" {
		t.Errorf("Unexpected quote data: %+v", q)
	}
	if len(s.quotes) != 1 {
		t.Errorf("Expected quote count 1, got %d", len(s.quotes))
	}
}

func TestGetQuotesByAuthor(t *testing.T) {
	s := NewService()
	s.AddQuote("Alice", "Q1")
	s.AddQuote("Bob", "Q2")
	s.AddQuote("Alice", "Q3")

	result := s.GetQuotesByAuthor("Alice")
	if len(result) != 2 {
		t.Errorf("Expected 2 quotes, got %d", len(result))
	}
}

func TestGetRandomQuote(t *testing.T) {
	s := NewService()
	s.AddQuote("A", "Q1")
	q := s.GetRandomQuote()
	if q == nil {
		t.Fatal("Expected random quote, got nil")
	}
}

func TestDeleteQuote(t *testing.T) {
	s := NewService()
	q := s.AddQuote("A", "Q1")

	deleted := s.DeleteQuote(q.ID)
	if !deleted {
		t.Error("Expected quote to be deleted")
	}
	if len(s.quotes) != 0 {
		t.Errorf("Expected no quotes left, got %d", len(s.quotes))
	}
}
