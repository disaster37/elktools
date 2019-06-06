// Code generated from specification version 7.0.0-SNAPSHOT: DO NOT EDIT

package xpack

import (
	"context"
	"strings"
)

func newIlmGetStatusFunc(t Transport) IlmGetStatus {
	return func(o ...func(*IlmGetStatusRequest)) (*Response, error) {
		var r = IlmGetStatusRequest{}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/current/ilm-get-status.html.
//
type IlmGetStatus func(o ...func(*IlmGetStatusRequest)) (*Response, error)

// IlmGetStatusRequest configures the Ilm  Get Status API request.
//
type IlmGetStatusRequest struct {
	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r IlmGetStatusRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "GET"

	path.Grow(len("/_ilm/status"))
	path.WriteString("/_ilm/status")

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
func (f IlmGetStatus) WithContext(v context.Context) func(*IlmGetStatusRequest) {
	return func(r *IlmGetStatusRequest) {
		r.ctx = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f IlmGetStatus) WithPretty() func(*IlmGetStatusRequest) {
	return func(r *IlmGetStatusRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f IlmGetStatus) WithHuman() func(*IlmGetStatusRequest) {
	return func(r *IlmGetStatusRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f IlmGetStatus) WithErrorTrace() func(*IlmGetStatusRequest) {
	return func(r *IlmGetStatusRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f IlmGetStatus) WithFilterPath(v ...string) func(*IlmGetStatusRequest) {
	return func(r *IlmGetStatusRequest) {
		r.FilterPath = v
	}
}
