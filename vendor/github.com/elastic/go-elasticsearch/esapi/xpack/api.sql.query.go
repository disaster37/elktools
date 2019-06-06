// Code generated from specification version 7.0.0-SNAPSHOT: DO NOT EDIT

package xpack

import (
	"context"
	"io"
	"strings"
)

func newSqlQueryFunc(t Transport) SqlQuery {
	return func(body io.Reader, o ...func(*SqlQueryRequest)) (*Response, error) {
		var r = SqlQueryRequest{Body: body}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

//
// See full documentation at Execute SQL.
//
type SqlQuery func(body io.Reader, o ...func(*SqlQueryRequest)) (*Response, error)

// SqlQueryRequest configures the Sql Query API request.
//
type SqlQueryRequest struct {
	Body io.Reader

	Format string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r SqlQueryRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "POST"

	path.Grow(len("/_sql"))
	path.WriteString("/_sql")

	params = make(map[string]string)

	if r.Format != "" {
		params["format"] = r.Format
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
func (f SqlQuery) WithContext(v context.Context) func(*SqlQueryRequest) {
	return func(r *SqlQueryRequest) {
		r.ctx = v
	}
}

// WithFormat - a short version of the accept header, e.g. json, yaml.
//
func (f SqlQuery) WithFormat(v string) func(*SqlQueryRequest) {
	return func(r *SqlQueryRequest) {
		r.Format = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f SqlQuery) WithPretty() func(*SqlQueryRequest) {
	return func(r *SqlQueryRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f SqlQuery) WithHuman() func(*SqlQueryRequest) {
	return func(r *SqlQueryRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f SqlQuery) WithErrorTrace() func(*SqlQueryRequest) {
	return func(r *SqlQueryRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f SqlQuery) WithFilterPath(v ...string) func(*SqlQueryRequest) {
	return func(r *SqlQueryRequest) {
		r.FilterPath = v
	}
}
