// Code generated from specification version 7.0.0-SNAPSHOT: DO NOT EDIT

package xpack

import (
	"context"
	"strings"
)

func newSecurityGetRoleFunc(t Transport) SecurityGetRole {
	return func(o ...func(*SecurityGetRoleRequest)) (*Response, error) {
		var r = SecurityGetRoleRequest{}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-get-role.html.
//
type SecurityGetRole func(o ...func(*SecurityGetRoleRequest)) (*Response, error)

// SecurityGetRoleRequest configures the Security  Get Role API request.
//
type SecurityGetRoleRequest struct {
	Name string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r SecurityGetRoleRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "GET"

	path.Grow(1 + len("_security") + 1 + len("role") + 1 + len(r.Name))
	path.WriteString("/")
	path.WriteString("_security")
	path.WriteString("/")
	path.WriteString("role")
	if r.Name != "" {
		path.WriteString("/")
		path.WriteString(r.Name)
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
func (f SecurityGetRole) WithContext(v context.Context) func(*SecurityGetRoleRequest) {
	return func(r *SecurityGetRoleRequest) {
		r.ctx = v
	}
}

// WithName - role name.
//
func (f SecurityGetRole) WithName(v string) func(*SecurityGetRoleRequest) {
	return func(r *SecurityGetRoleRequest) {
		r.Name = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f SecurityGetRole) WithPretty() func(*SecurityGetRoleRequest) {
	return func(r *SecurityGetRoleRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f SecurityGetRole) WithHuman() func(*SecurityGetRoleRequest) {
	return func(r *SecurityGetRoleRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f SecurityGetRole) WithErrorTrace() func(*SecurityGetRoleRequest) {
	return func(r *SecurityGetRoleRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f SecurityGetRole) WithFilterPath(v ...string) func(*SecurityGetRoleRequest) {
	return func(r *SecurityGetRoleRequest) {
		r.FilterPath = v
	}
}
