package gophy

import (
	"fmt"

	"github.com/monkeydioude/tools/http"
)

const (
	//APIHost is the base API Host
	APIHost = "https://api.giphy.com"
)

// Request is the interface that should be used for requesting Giphy
type Request interface {
	Build() (string, error)
	GetMethod() string
	GetPath() string
}

// Gophy is the core of this lib. Handles request and responses from giphy
type Gophy struct {
	APIKey  string
	APIHost string
	Client  http.Client
}

// NewGophy returns a pointer to Gophy using the client API Key
func NewGophy(APIKey string) *Gophy {
	return &Gophy{
		APIKey:  APIKey,
		APIHost: APIHost,
		Client:  http.DefaultClient,
	}
}

// Request handles any request to the Giphy endpoint
func (g *Gophy) Request(r Request) ([]byte, error) {
	query, err := r.Build()

	if err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf("%s/%s%s&api_key=%s", g.APIHost, r.GetPath(), query, g.APIKey)
	res, err := http.Request(nil, nil, endpoint, r.GetMethod())

	if err != nil {
		return nil, err
	}

	return http.NewBytesResponseHTTP(res)
}
