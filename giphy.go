package gophy

import (
	"github.com/monkeydioude/tools/http"
)

const (
	//APIHost is the base API Host
	APIHost = "api.giphy.com"
)

// Request is the interface that should be used for requesting Giphy
type Request interface {
	Do() (string, error)
	GetMethod() string
	GetPath() string
}

// Gophy is the core of this lib. Handles request and responses from giphy
type Gophy struct {
	APIKey  string
	APIHost string
}

// NewGophy returns a pointer to Gophy using the client API Key
func NewGophy(APIKey string) *Gophy {
	return &Gophy{
		APIKey:  APIKey,
		APIHost: APIHost,
	}
}

// Request handles any request to the Giphy endpoint
func (g *Gophy) Request(r *Request) error {
	http.Request()
	return nil
}
