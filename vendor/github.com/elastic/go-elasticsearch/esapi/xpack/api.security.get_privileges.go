// Code generated from specification version 7.0.0-SNAPSHOT: DO NOT EDIT

package xpack

import (
	"context"
	"strings"
)

func newSecurityGetPrivilegesFunc(t Transport) SecurityGetPrivileges {
	return func(o ...func(*SecurityGetPrivilegesRequest)) (*Response, error) {
		var r = SecurityGetPrivilegesRequest{}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

//
// See full documentation at TODO.
//
type SecurityGetPrivileges func(o ...func(*SecurityGetPrivilegesRequest)) (*Response, error)

// SecurityGetPrivilegesRequest configures the Security  Get Privileges API request.
//
type SecurityGetPrivilegesRequest struct {
	Application string
	Name        string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r SecurityGetPrivilegesRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "GET"

	path.Grow(1 + len("_security") + 1 + len("privilege") + 1 + len(r.Application) + 1 + len(r.Name))
	path.WriteString("/")
	path.WriteString("_security")
	path.WriteString("/")
	path.WriteString("privilege")
	if r.Application != "" {
		path.WriteString("/")
		path.WriteString(r.Application)
	}
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
func (f SecurityGetPrivileges) WithContext(v context.Context) func(*SecurityGetPrivilegesRequest) {
	return func(r *SecurityGetPrivilegesRequest) {
		r.ctx = v
	}
}

// WithApplication - application name.
//
func (f SecurityGetPrivileges) WithApplication(v string) func(*SecurityGetPrivilegesRequest) {
	return func(r *SecurityGetPrivilegesRequest) {
		r.Application = v
	}
}

// WithName - privilege name.
//
func (f SecurityGetPrivileges) WithName(v string) func(*SecurityGetPrivilegesRequest) {
	return func(r *SecurityGetPrivilegesRequest) {
		r.Name = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f SecurityGetPrivileges) WithPretty() func(*SecurityGetPrivilegesRequest) {
	return func(r *SecurityGetPrivilegesRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f SecurityGetPrivileges) WithHuman() func(*SecurityGetPrivilegesRequest) {
	return func(r *SecurityGetPrivilegesRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f SecurityGetPrivileges) WithErrorTrace() func(*SecurityGetPrivilegesRequest) {
	return func(r *SecurityGetPrivilegesRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f SecurityGetPrivileges) WithFilterPath(v ...string) func(*SecurityGetPrivilegesRequest) {
	return func(r *SecurityGetPrivilegesRequest) {
		r.FilterPath = v
	}
}
