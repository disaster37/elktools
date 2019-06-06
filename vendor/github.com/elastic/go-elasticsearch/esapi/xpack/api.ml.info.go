// Code generated from specification version 7.0.0-SNAPSHOT: DO NOT EDIT

package xpack

import (
	"context"
	"strings"
)

func newMlInfoFunc(t Transport) MlInfo {
	return func(o ...func(*MlInfoRequest)) (*Response, error) {
		var r = MlInfoRequest{}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

//
type MlInfo func(o ...func(*MlInfoRequest)) (*Response, error)

// MlInfoRequest configures the Ml Info API request.
//
type MlInfoRequest struct {
	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r MlInfoRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "GET"

	path.Grow(len("/_ml/info"))
	path.WriteString("/_ml/info")

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
func (f MlInfo) WithContext(v context.Context) func(*MlInfoRequest) {
	return func(r *MlInfoRequest) {
		r.ctx = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f MlInfo) WithPretty() func(*MlInfoRequest) {
	return func(r *MlInfoRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f MlInfo) WithHuman() func(*MlInfoRequest) {
	return func(r *MlInfoRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f MlInfo) WithErrorTrace() func(*MlInfoRequest) {
	return func(r *MlInfoRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f MlInfo) WithFilterPath(v ...string) func(*MlInfoRequest) {
	return func(r *MlInfoRequest) {
		r.FilterPath = v
	}
}
