package request

import "testing"

func TestIGetSearchQuery(t *testing.T) {
	s := NewSearch("nicki+minaj")
	goal := "?q=nicki+minaj&offset=0"

	if trial, _ := s.Build(); trial != goal {
		t.Errorf("\nExpected %s \nGot %s", goal, trial)
	}
}

func TestIFailOnEmptyQuery(t *testing.T) {
	s := &Search{}
	if _, err := s.Build(); err == nil {
		t.Error("Error should have been returned")
	}
}

func TestICanGetSearchWithAllParaneters(t *testing.T) {
	s := &Search{
		Query:  "freddy+mercury",
		Offset: 1,
		Limit:  10,
		Rating: "m",
		Lang:   "kl",
		Fmt:    "json",
	}

	goal := "?q=freddy+mercury&offset=1&limit=10&rating=m&lang=kl&fmt=json"

	if trial, _ := s.Build(); trial != goal {
		t.Errorf("\nExp %s \nGot %s", goal, trial)
	}
}
