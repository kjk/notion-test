package notion

// tries to mimic https://github.com/makenotion/notion-sdk-js/blob/main/src/Client.ts

import (
	"io"
	"net/http"
	"time"
)

const DefaultNotionVersion = "2021-05-13"

type ClientOptions struct {
	Auth          string
	Timeout       time.Duration
	BaseURL       string
	Logger        io.Writer
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
	version    string
	Auth       string
	HTTPClient *http.Client
}

func NewClient(opts ClientOptions) (*Client, error) {
	res := &Client{
		version: DefaultNotionVersion,
	}
	if opts.NotionVersion != "" {
		res.version = opts.NotionVersion
	}
	res.Auth = opts.Auth

	return res, nil
}

func (c *Client) BlocksList(args *BlocksChildrenListParameters) (*BlocksChildrenListResponse, error) {

	return nil, nil
}
