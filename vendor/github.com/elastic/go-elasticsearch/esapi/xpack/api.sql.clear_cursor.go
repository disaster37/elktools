// Code generated from specification version 7.0.0-SNAPSHOT: DO NOT EDIT

package xpack

import (
	"context"
	"io"
	"strings"
)

func newSqlClearCursorFunc(t Transport) SqlClearCursor {
	return func(body io.Reader, o ...func(*SqlClearCursorRequest)) (*Response, error) {
		var r = SqlClearCursorRequest{Body: body}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

//
// See full documentation at Clear SQL cursor.
//
type SqlClearCursor func(body io.Reader, o ...func(*SqlClearCursorRequest)) (*Response, error)

// SqlClearCursorRequest configures the Sql  Clear Cursor API request.
//
type SqlClearCursorRequest struct {
	Body io.Reader

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r SqlClearCursorRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "POST"

	path.Grow(len("/_sql/close"))
	path.WriteString("/_sql/close")

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
func (f SqlClearCursor) WithContext(v context.Context) func(*SqlClearCursorRequest) {
	return func(r *SqlClearCursorRequest) {
		r.ctx = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f SqlClearCursor) WithPretty() func(*SqlClearCursorRequest) {
	return func(r *SqlClearCursorRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f SqlClearCursor) WithHuman() func(*SqlClearCursorRequest) {
	return func(r *SqlClearCursorRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f SqlClearCursor) WithErrorTrace() func(*SqlClearCursorRequest) {
	return func(r *SqlClearCursorRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f SqlClearCursor) WithFilterPath(v ...string) func(*SqlClearCursorRequest) {
	return func(r *SqlClearCursorRequest) {
		r.FilterPath = v
	}
}
