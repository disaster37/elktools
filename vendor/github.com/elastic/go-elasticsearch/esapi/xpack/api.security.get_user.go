// Code generated from specification version 7.0.0-SNAPSHOT: DO NOT EDIT

package xpack

import (
	"context"
	"strings"
)

func newSecurityGetUserFunc(t Transport) SecurityGetUser {
	return func(o ...func(*SecurityGetUserRequest)) (*Response, error) {
		var r = SecurityGetUserRequest{}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-get-user.html.
//
type SecurityGetUser func(o ...func(*SecurityGetUserRequest)) (*Response, error)

// SecurityGetUserRequest configures the Security  Get User API request.
//
type SecurityGetUserRequest struct {
	Username []string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r SecurityGetUserRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "GET"

	path.Grow(1 + len("_security") + 1 + len("user") + 1 + len(strings.Join(r.Username, ",")))
	path.WriteString("/")
	path.WriteString("_security")
	path.WriteString("/")
	path.WriteString("user")
	if len(r.Username) > 0 {
		path.WriteString("/")
		path.WriteString(strings.Join(r.Username, ","))
	}

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
func (f SecurityGetUser) WithContext(v context.Context) func(*SecurityGetUserRequest) {
	return func(r *SecurityGetUserRequest) {
		r.ctx = v
	}
}

// WithUsername - a list of usernames.
//
func (f SecurityGetUser) WithUsername(v ...string) func(*SecurityGetUserRequest) {
	return func(r *SecurityGetUserRequest) {
		r.Username = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f SecurityGetUser) WithPretty() func(*SecurityGetUserRequest) {
	return func(r *SecurityGetUserRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f SecurityGetUser) WithHuman() func(*SecurityGetUserRequest) {
	return func(r *SecurityGetUserRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f SecurityGetUser) WithErrorTrace() func(*SecurityGetUserRequest) {
	return func(r *SecurityGetUserRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f SecurityGetUser) WithFilterPath(v ...string) func(*SecurityGetUserRequest) {
	return func(r *SecurityGetUserRequest) {
		r.FilterPath = v
	}
}
