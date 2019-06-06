// Code generated from specification version 7.0.0-SNAPSHOT: DO NOT EDIT

package xpack

import (
	"context"
	"strings"
)

func newIlmStopFunc(t Transport) IlmStop {
	return func(o ...func(*IlmStopRequest)) (*Response, error) {
		var r = IlmStopRequest{}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/current/ilm-stop.html.
//
type IlmStop func(o ...func(*IlmStopRequest)) (*Response, error)

// IlmStopRequest configures the Ilm Stop API request.
//
type IlmStopRequest struct {
	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r IlmStopRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "POST"

	path.Grow(len("/_ilm/stop"))
	path.WriteString("/_ilm/stop")

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
func (f IlmStop) WithContext(v context.Context) func(*IlmStopRequest) {
	return func(r *IlmStopRequest) {
		r.ctx = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f IlmStop) WithPretty() func(*IlmStopRequest) {
	return func(r *IlmStopRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f IlmStop) WithHuman() func(*IlmStopRequest) {
	return func(r *IlmStopRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f IlmStop) WithErrorTrace() func(*IlmStopRequest) {
	return func(r *IlmStopRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f IlmStop) WithFilterPath(v ...string) func(*IlmStopRequest) {
	return func(r *IlmStopRequest) {
		r.FilterPath = v
	}
}
