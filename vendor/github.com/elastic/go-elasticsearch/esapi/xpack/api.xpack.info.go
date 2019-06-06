// Code generated from specification version 7.0.0-SNAPSHOT: DO NOT EDIT

package xpack

import (
	"context"
	"strings"
)

func newXpackInfoFunc(t Transport) XpackInfo {
	return func(o ...func(*XpackInfoRequest)) (*Response, error) {
		var r = XpackInfoRequest{}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/current/info-api.html.
//
type XpackInfo func(o ...func(*XpackInfoRequest)) (*Response, error)

// XpackInfoRequest configures the Xpack Info API request.
//
type XpackInfoRequest struct {
	Categories []string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r XpackInfoRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "GET"

	path.Grow(len("/_xpack"))
	path.WriteString("/_xpack")

	params = make(map[string]string)

	if len(r.Categories) > 0 {
		params["categories"] = strings.Join(r.Categories, ",")
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
func (f XpackInfo) WithContext(v context.Context) func(*XpackInfoRequest) {
	return func(r *XpackInfoRequest) {
		r.ctx = v
	}
}

// WithCategories - comma-separated list of info categories. can be any of: build, license, features.
//
func (f XpackInfo) WithCategories(v ...string) func(*XpackInfoRequest) {
	return func(r *XpackInfoRequest) {
		r.Categories = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f XpackInfo) WithPretty() func(*XpackInfoRequest) {
	return func(r *XpackInfoRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f XpackInfo) WithHuman() func(*XpackInfoRequest) {
	return func(r *XpackInfoRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f XpackInfo) WithErrorTrace() func(*XpackInfoRequest) {
	return func(r *XpackInfoRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f XpackInfo) WithFilterPath(v ...string) func(*XpackInfoRequest) {
	return func(r *XpackInfoRequest) {
		r.FilterPath = v
	}
}
