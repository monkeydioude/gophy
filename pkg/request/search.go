package request

import (
	"errors"
	"strconv"
	"strings"
)

type Search struct {
	Query  string
	Limit  int
	Offset int
	Rating string
	Lang   string
	Fmt    string
}

func (s *Search) Build() (string, error) {
	if s.Query == "" {
		return "", errors.New("At least 'query' parameter should be provided")
	}
	var q strings.Builder

	q.WriteString("?q=")
	q.WriteString(s.Query)
	q.WriteString("&offset=")
	q.WriteString(strconv.Itoa(s.Offset))

	if s.Limit != 0 {
		q.WriteString("&limit=")
		q.WriteString(strconv.Itoa(s.Limit))
	}

	if s.Rating != "" {
		q.WriteString("&rating=")
		q.WriteString(s.Rating)
	}

	if s.Lang != "" {
		q.WriteString("&lang=")
		q.WriteString(s.Lang)
	}

	if s.Fmt != "" {
		q.WriteString("&fmt=")
		q.WriteString(s.Fmt)
	}

	return q.String(), nil
}

func (s *Search) GetMethod() string {
	return "GET"
}

func (s *Search) GetPath() string {
	return "v1/gifs/search"
}

// NewSearch returns a Search from a query
func NewSearch(query string) *Search {
	return &Search{
		Query: query,
	}
}
