package rest

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"
)

type Doer interface {
	Do(r *http.Request) (*http.Response, error)
}

type DefaultHTTPClient struct {
	doer          Doer
	userAgent     string
	authorization string
}

func (c *DefaultHTTPClient) Request(req *request) (*DiscordResponse, error) {
	var reader io.Reader = nil
	if req.body != nil {
		reader = bytes.NewReader(req.body)
	}

	var rawReq *http.Request
	var err error

	if req.ctx != nil {
		rawReq, err = http.NewRequestWithContext(req.ctx, req.method, req.path, reader)
	} else {
		rawReq, err = http.NewRequest(req.method, req.path, reader)
	}
	if err != nil {
		return nil, err
	}

	if req.headers != nil {
		rawReq.Header = req.headers.Clone()
	}

	if reader != nil {
		rawReq.Header.Set("Content-Type", req.contentType)
	}

	if c.userAgent != "" {
		rawReq.Header.Set("User-Agent", c.userAgent)
	}

	if !req.omitAuth {
		rawReq.Header.Set("Authorization", c.authorization)
	}

	resp, err := c.doer.Do(rawReq)
	if err != nil {
		return nil, err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return &DiscordResponse{
		Body:       respBody,
		StatusCode: resp.StatusCode,
		Header:     resp.Header,
	}, nil
}

// TestHTTPClient is a replacement HTTP client that can be used during testing.
type TestHTTPClient struct {
	T               *testing.T
	ExpectedRequest *request
	Response        *DiscordResponse
	Error           error
}

func (c *TestHTTPClient) Request(req *request) (*DiscordResponse, error) {
	if !reflect.DeepEqual(req, c.ExpectedRequest) && c.T != nil {
		c.T.Errorf("Request does not match expected request")
	}

	return c.Response, c.Error
}
