// Code generated from specification version 7.0.0-SNAPSHOT: DO NOT EDIT

package xpack

import (
	"context"
	"strings"
)

func newSecurityDeleteUserFunc(t Transport) SecurityDeleteUser {
	return func(username string, o ...func(*SecurityDeleteUserRequest)) (*Response, error) {
		var r = SecurityDeleteUserRequest{Username: username}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-delete-user.html.
//
type SecurityDeleteUser func(username string, o ...func(*SecurityDeleteUserRequest)) (*Response, error)

// SecurityDeleteUserRequest configures the Security  Delete User API request.
//
type SecurityDeleteUserRequest struct {
	Username string
	Refresh  string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r SecurityDeleteUserRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "DELETE"

	path.Grow(1 + len("_security") + 1 + len("user") + 1 + len(r.Username))
	path.WriteString("/")
	path.WriteString("_security")
	path.WriteString("/")
	path.WriteString("user")
	path.WriteString("/")
	path.WriteString(r.Username)

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
func (f SecurityDeleteUser) WithContext(v context.Context) func(*SecurityDeleteUserRequest) {
	return func(r *SecurityDeleteUserRequest) {
		r.ctx = v
	}
}

// WithRefresh - if `true` (the default) then refresh the affected shards to make this operation visible to search, if `wait_for` then wait for a refresh to make this operation visible to search, if `false` then do nothing with refreshes..
//
func (f SecurityDeleteUser) WithRefresh(v string) func(*SecurityDeleteUserRequest) {
	return func(r *SecurityDeleteUserRequest) {
		r.Refresh = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f SecurityDeleteUser) WithPretty() func(*SecurityDeleteUserRequest) {
	return func(r *SecurityDeleteUserRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f SecurityDeleteUser) WithHuman() func(*SecurityDeleteUserRequest) {
	return func(r *SecurityDeleteUserRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f SecurityDeleteUser) WithErrorTrace() func(*SecurityDeleteUserRequest) {
	return func(r *SecurityDeleteUserRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f SecurityDeleteUser) WithFilterPath(v ...string) func(*SecurityDeleteUserRequest) {
	return func(r *SecurityDeleteUserRequest) {
		r.FilterPath = v
	}
}
