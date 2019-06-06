// Code generated from specification version 7.0.0-SNAPSHOT: DO NOT EDIT

package xpack

import (
	"context"
	"strings"
)

func newSecurityGetUserPrivilegesFunc(t Transport) SecurityGetUserPrivileges {
	return func(o ...func(*SecurityGetUserPrivilegesRequest)) (*Response, error) {
		var r = SecurityGetUserPrivilegesRequest{}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-get-user-privileges.html.
//
type SecurityGetUserPrivileges func(o ...func(*SecurityGetUserPrivilegesRequest)) (*Response, error)

// SecurityGetUserPrivilegesRequest configures the Security   Get User Privileges API request.
//
type SecurityGetUserPrivilegesRequest struct {
	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r SecurityGetUserPrivilegesRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "GET"

	path.Grow(len("/_security/user/_privileges"))
	path.WriteString("/_security/user/_privileges")

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
func (f SecurityGetUserPrivileges) WithContext(v context.Context) func(*SecurityGetUserPrivilegesRequest) {
	return func(r *SecurityGetUserPrivilegesRequest) {
		r.ctx = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f SecurityGetUserPrivileges) WithPretty() func(*SecurityGetUserPrivilegesRequest) {
	return func(r *SecurityGetUserPrivilegesRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f SecurityGetUserPrivileges) WithHuman() func(*SecurityGetUserPrivilegesRequest) {
	return func(r *SecurityGetUserPrivilegesRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f SecurityGetUserPrivileges) WithErrorTrace() func(*SecurityGetUserPrivilegesRequest) {
	return func(r *SecurityGetUserPrivilegesRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f SecurityGetUserPrivileges) WithFilterPath(v ...string) func(*SecurityGetUserPrivilegesRequest) {
	return func(r *SecurityGetUserPrivilegesRequest) {
		r.FilterPath = v
	}
}
