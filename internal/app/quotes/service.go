package quotes

import (
	"math/rand"
	"slices"
)

type Service struct {
	quotes []Quote
	nextID int
}

func NewService() *Service {
	return &Service{
		quotes: make([]Quote, 0),
		nextID: 1,
	}
}

func (s *Service) AddQuote(author, quote string) Quote {
	q := Quote{ID: s.nextID, Author: author, Quote: quote}

	s.quotes = append(s.quotes, q)
	s.nextID++

	return q
}

func (s *Service) GetAllQuotes() []Quote {
	result := make([]Quote, len(s.quotes))
	copy(result, s.quotes)

	return result
}

func (s *Service) GetQuotesByAuthor(author string) []Quote {
	var result []Quote

	for _, q := range s.quotes {
		if q.Author == author {
			result = append(result, q)
		}
	}

	return result
}

func (s *Service) GetRandomQuote() *Quote {
	if len(s.quotes) == 0 {
		return nil
	}

	q := s.quotes[rand.Intn(len(s.quotes))]

	return &q
}

func (s *Service) DeleteQuote(id int) bool {
	for i, q := range s.quotes {
		if id == q.ID {
			s.quotes = slices.Delete(s.quotes, i, i+1)
			s.nextID--
			return true
		}
	}

	return false
}
