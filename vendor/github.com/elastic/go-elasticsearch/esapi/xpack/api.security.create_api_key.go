// Code generated from specification version 7.0.0-SNAPSHOT: DO NOT EDIT

package xpack

import (
	"context"
	"io"
	"strings"
)

func newSecurityCreateApiKeyFunc(t Transport) SecurityCreateApiKey {
	return func(body io.Reader, o ...func(*SecurityCreateApiKeyRequest)) (*Response, error) {
		var r = SecurityCreateApiKeyRequest{Body: body}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-create-api-key.html.
//
type SecurityCreateApiKey func(body io.Reader, o ...func(*SecurityCreateApiKeyRequest)) (*Response, error)

// SecurityCreateApiKeyRequest configures the Security   Create Api Key API request.
//
type SecurityCreateApiKeyRequest struct {
	Body io.Reader

	Refresh string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r SecurityCreateApiKeyRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "PUT"

	path.Grow(len("/_security/api_key"))
	path.WriteString("/_security/api_key")

	params = make(map[string]string)

	if r.Refresh != "" {
		params["refresh"] = r.Refresh
	}

	if r.Pretty {
		params["pretty"] = "true"
	}

	if r.Human {
		params["human"] = "true"
	}

	if r.ErrorTrace {
		params["error_trace"] = "true"
	}

	if len(r.FilterPath) > 0 {
		params["filter_path"] = strings.Join(r.FilterPath, ",")
	}

	req, _ := newRequest(method, path.String(), r.Body)

	if len(params) > 0 {
		q := req.URL.Query()
		for k, v := range params {
			q.Set(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	if r.Body != nil {
		req.Header[headerContentType] = headerContentTypeJSON
	}

	if ctx != nil {
		req = req.WithContext(ctx)
	}

	res, err := transport.Perform(req)
	if err != nil {
		return nil, err
	}

	response := Response{
		StatusCode: res.StatusCode,
		Body:       res.Body,
		Header:     res.Header,
	}

	return &response, nil
}

// WithContext sets the request context.
//
func (f SecurityCreateApiKey) WithContext(v context.Context) func(*SecurityCreateApiKeyRequest) {
	return func(r *SecurityCreateApiKeyRequest) {
		r.ctx = v
	}
}

// WithRefresh - if `true` (the default) then refresh the affected shards to make this operation visible to search, if `wait_for` then wait for a refresh to make this operation visible to search, if `false` then do nothing with refreshes..
//
func (f SecurityCreateApiKey) WithRefresh(v string) func(*SecurityCreateApiKeyRequest) {
	return func(r *SecurityCreateApiKeyRequest) {
		r.Refresh = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f SecurityCreateApiKey) WithPretty() func(*SecurityCreateApiKeyRequest) {
	return func(r *SecurityCreateApiKeyRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f SecurityCreateApiKey) WithHuman() func(*SecurityCreateApiKeyRequest) {
	return func(r *SecurityCreateApiKeyRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f SecurityCreateApiKey) WithErrorTrace() func(*SecurityCreateApiKeyRequest) {
	return func(r *SecurityCreateApiKeyRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f SecurityCreateApiKey) WithFilterPath(v ...string) func(*SecurityCreateApiKeyRequest) {
	return func(r *SecurityCreateApiKeyRequest) {
		r.FilterPath = v
	}
}
