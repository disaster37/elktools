// Code generated from specification version 7.0.0-SNAPSHOT: DO NOT EDIT

package xpack

import (
	"context"
	"io"
	"strings"
)

func newSecurityPutRoleFunc(t Transport) SecurityPutRole {
	return func(body io.Reader, name string, o ...func(*SecurityPutRoleRequest)) (*Response, error) {
		var r = SecurityPutRoleRequest{Body: body, Name: name}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-put-role.html.
//
type SecurityPutRole func(body io.Reader, name string, o ...func(*SecurityPutRoleRequest)) (*Response, error)

// SecurityPutRoleRequest configures the Security  Put Role API request.
//
type SecurityPutRoleRequest struct {
	Body io.Reader

	Name    string
	Refresh string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r SecurityPutRoleRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "PUT"

	path.Grow(1 + len("_security") + 1 + len("role") + 1 + len(r.Name))
	path.WriteString("/")
	path.WriteString("_security")
	path.WriteString("/")
	path.WriteString("role")
	path.WriteString("/")
	path.WriteString(r.Name)

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
func (f SecurityPutRole) WithContext(v context.Context) func(*SecurityPutRoleRequest) {
	return func(r *SecurityPutRoleRequest) {
		r.ctx = v
	}
}

// WithRefresh - if `true` (the default) then refresh the affected shards to make this operation visible to search, if `wait_for` then wait for a refresh to make this operation visible to search, if `false` then do nothing with refreshes..
//
func (f SecurityPutRole) WithRefresh(v string) func(*SecurityPutRoleRequest) {
	return func(r *SecurityPutRoleRequest) {
		r.Refresh = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f SecurityPutRole) WithPretty() func(*SecurityPutRoleRequest) {
	return func(r *SecurityPutRoleRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f SecurityPutRole) WithHuman() func(*SecurityPutRoleRequest) {
	return func(r *SecurityPutRoleRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f SecurityPutRole) WithErrorTrace() func(*SecurityPutRoleRequest) {
	return func(r *SecurityPutRoleRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f SecurityPutRole) WithFilterPath(v ...string) func(*SecurityPutRoleRequest) {
	return func(r *SecurityPutRoleRequest) {
		r.FilterPath = v
	}
}
