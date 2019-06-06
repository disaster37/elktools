// Code generated from specification version 7.0.0-SNAPSHOT: DO NOT EDIT

package xpack

import (
	"context"
	"strings"
)

func newSecurityAuthenticateFunc(t Transport) SecurityAuthenticate {
	return func(o ...func(*SecurityAuthenticateRequest)) (*Response, error) {
		var r = SecurityAuthenticateRequest{}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-authenticate.html.
//
type SecurityAuthenticate func(o ...func(*SecurityAuthenticateRequest)) (*Response, error)

// SecurityAuthenticateRequest configures the Security Authenticate API request.
//
type SecurityAuthenticateRequest struct {
	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r SecurityAuthenticateRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "GET"

	path.Grow(len("/_security/_authenticate"))
	path.WriteString("/_security/_authenticate")

	params = make(map[string]string)

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

	req, _ := newRequest(method, path.String(), nil)

	if len(params) > 0 {
		q := req.URL.Query()
		for k, v := range params {
			q.Set(k, v)
		}
		req.URL.RawQuery = q.Encode()
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
func (f SecurityAuthenticate) WithContext(v context.Context) func(*SecurityAuthenticateRequest) {
	return func(r *SecurityAuthenticateRequest) {
		r.ctx = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f SecurityAuthenticate) WithPretty() func(*SecurityAuthenticateRequest) {
	return func(r *SecurityAuthenticateRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f SecurityAuthenticate) WithHuman() func(*SecurityAuthenticateRequest) {
	return func(r *SecurityAuthenticateRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f SecurityAuthenticate) WithErrorTrace() func(*SecurityAuthenticateRequest) {
	return func(r *SecurityAuthenticateRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f SecurityAuthenticate) WithFilterPath(v ...string) func(*SecurityAuthenticateRequest) {
	return func(r *SecurityAuthenticateRequest) {
		r.FilterPath = v
	}
}
