// Code generated from specification version 7.0.0-SNAPSHOT: DO NOT EDIT

package xpack

import (
	"context"
	"strings"
)

func newIlmStartFunc(t Transport) IlmStart {
	return func(o ...func(*IlmStartRequest)) (*Response, error) {
		var r = IlmStartRequest{}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/current/ilm-start.html.
//
type IlmStart func(o ...func(*IlmStartRequest)) (*Response, error)

// IlmStartRequest configures the Ilm Start API request.
//
type IlmStartRequest struct {
	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r IlmStartRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "POST"

	path.Grow(len("/_ilm/start"))
	path.WriteString("/_ilm/start")

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
func (f IlmStart) WithContext(v context.Context) func(*IlmStartRequest) {
	return func(r *IlmStartRequest) {
		r.ctx = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f IlmStart) WithPretty() func(*IlmStartRequest) {
	return func(r *IlmStartRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f IlmStart) WithHuman() func(*IlmStartRequest) {
	return func(r *IlmStartRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f IlmStart) WithErrorTrace() func(*IlmStartRequest) {
	return func(r *IlmStartRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f IlmStart) WithFilterPath(v ...string) func(*IlmStartRequest) {
	return func(r *IlmStartRequest) {
		r.FilterPath = v
	}
}
