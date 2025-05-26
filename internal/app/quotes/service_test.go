package quotes

import (
	"testing"
)

func TestAddQuote(t *testing.T) {
	t.Log("Creating a new quote service")
	s := NewService()

	t.Log("Adding a quote by Author1")
	q := s.AddQuote("Author1", "Quote1")

	t.Logf("Added quote: %+v", q)

	if q.ID != 1 {
		t.Errorf("Expected ID 1, got %d", q.ID)
	}
}

func TestGetQuotesByAuthor(t *testing.T) {
	t.Log("Creating service and adding quotes")
	s := NewService()
	s.AddQuote("Alice", "Q1")
	s.AddQuote("Bob", "Q2")
	s.AddQuote("Alice", "Q3")

	t.Log("Fetching quotes by Alice")
	result := s.GetQuotesByAuthor("Alice")

	t.Logf("Found %d quote(s): %+v", len(result), result)
	if len(result) != 2 {
		t.Errorf("Expected 2 quotes, got %d", len(result))
	}
}

func TestGetRandomQuote(t *testing.T) {
	t.Log("Creating service and adding one quote")
	s := NewService()
	s.AddQuote("A", "Q1")

	t.Log("Fetching random quote")
	q := s.GetRandomQuote()
	t.Logf("Random quote: %+v", q)

	if q == nil {
		t.Fatal("Expected non-nil quote")
	}
}

func TestDeleteQuote(t *testing.T) {
	t.Log("Creating service and adding quote")
	s := NewService()
	q := s.AddQuote("A", "Q1")

	t.Logf("Deleting quote with ID: %d", q.ID)
	deleted := s.DeleteQuote(q.ID)

	t.Logf("Delete success: %v", deleted)
	if !deleted {
		t.Error("Expected quote to be deleted")
	}

	if len(s.quotes) != 0 {
		t.Errorf("Expected 0 quotes left, got %d", len(s.quotes))
	}
}
