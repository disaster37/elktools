// Code generated from specification version 7.0.0-SNAPSHOT: DO NOT EDIT

package xpack

import (
	"context"
	"strings"
	"time"
)

func newXpackUsageFunc(t Transport) XpackUsage {
	return func(o ...func(*XpackUsageRequest)) (*Response, error) {
		var r = XpackUsageRequest{}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

//
// See full documentation at Retrieve information about xpack features usage.
//
type XpackUsage func(o ...func(*XpackUsageRequest)) (*Response, error)

// XpackUsageRequest configures the Xpack Usage API request.
//
type XpackUsageRequest struct {
	MasterTimeout time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r XpackUsageRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "GET"

	path.Grow(len("/_xpack/usage"))
	path.WriteString("/_xpack/usage")

	params = make(map[string]string)

	if r.MasterTimeout != 0 {
		params["master_timeout"] = formatDuration(r.MasterTimeout)
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
func (f XpackUsage) WithContext(v context.Context) func(*XpackUsageRequest) {
	return func(r *XpackUsageRequest) {
		r.ctx = v
	}
}

// WithMasterTimeout - specify timeout for watch write operation.
//
func (f XpackUsage) WithMasterTimeout(v time.Duration) func(*XpackUsageRequest) {
	return func(r *XpackUsageRequest) {
		r.MasterTimeout = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f XpackUsage) WithPretty() func(*XpackUsageRequest) {
	return func(r *XpackUsageRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f XpackUsage) WithHuman() func(*XpackUsageRequest) {
	return func(r *XpackUsageRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f XpackUsage) WithErrorTrace() func(*XpackUsageRequest) {
	return func(r *XpackUsageRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f XpackUsage) WithFilterPath(v ...string) func(*XpackUsageRequest) {
	return func(r *XpackUsageRequest) {
		r.FilterPath = v
	}
}
