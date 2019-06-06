// Code generated from specification version 7.0.0-SNAPSHOT: DO NOT EDIT

package xpack

import (
	"context"
	"io"
	"strings"
)

func newCcrForgetFollowerFunc(t Transport) CcrForgetFollower {
	return func(index string, body io.Reader, o ...func(*CcrForgetFollowerRequest)) (*Response, error) {
		var r = CcrForgetFollowerRequest{Index: index, Body: body}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

//
// See full documentation at http://www.elastic.co/guide/en/elasticsearch/reference/current.
//
type CcrForgetFollower func(index string, body io.Reader, o ...func(*CcrForgetFollowerRequest)) (*Response, error)

// CcrForgetFollowerRequest configures the Ccr  Forget Follower API request.
//
type CcrForgetFollowerRequest struct {
	Index string
	Body  io.Reader

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r CcrForgetFollowerRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "POST"

	path.Grow(1 + len(r.Index) + 1 + len("_ccr") + 1 + len("forget_follower"))
	path.WriteString("/")
	path.WriteString(r.Index)
	path.WriteString("/")
	path.WriteString("_ccr")
	path.WriteString("/")
	path.WriteString("forget_follower")

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
func (f CcrForgetFollower) WithContext(v context.Context) func(*CcrForgetFollowerRequest) {
	return func(r *CcrForgetFollowerRequest) {
		r.ctx = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f CcrForgetFollower) WithPretty() func(*CcrForgetFollowerRequest) {
	return func(r *CcrForgetFollowerRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f CcrForgetFollower) WithHuman() func(*CcrForgetFollowerRequest) {
	return func(r *CcrForgetFollowerRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f CcrForgetFollower) WithErrorTrace() func(*CcrForgetFollowerRequest) {
	return func(r *CcrForgetFollowerRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f CcrForgetFollower) WithFilterPath(v ...string) func(*CcrForgetFollowerRequest) {
	return func(r *CcrForgetFollowerRequest) {
		r.FilterPath = v
	}
}
