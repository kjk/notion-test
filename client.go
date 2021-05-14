package notion

// tries to mimic https://github.com/makenotion/notion-sdk-js/blob/main/src/Client.ts

import (
	"errors"
	"net/http"
	"time"
)

const DefaultNotionVersion = "2021-05-13"

type Logger interface {
	Log(s string)
}

type ClientOptions struct {
	Auth          string
	Timeout       time.Time
	BaseURL       string
	Logger        *Logger
	Agent         string
	NotionVersion string
}

/*
export interface RequestParameters {
  path: string;
  method: Method;
  query?: QueryParams;
  body?: Record<string, unknown>;
  auth?: string;
}
*/

type RequestParameters struct {
	Path   string
	Method string // TODO: better method
	// query?: QueryParams;
	// body?: Record<string, unknown>;
	Auth string
}

type Client struct {
	version string
	Auth    string
	Http    *http.Client
}

func NewClient(auth string) (*Client, error) {
	if auth == "" {
		return nil, errors.New("needs a valid token")
	}
	res := &Client{
		version: VERSION,
		Auth:    auth,
	}
	return res, nil
}
