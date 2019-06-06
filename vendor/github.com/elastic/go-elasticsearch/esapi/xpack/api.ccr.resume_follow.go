// Code generated from specification version 7.0.0-SNAPSHOT: DO NOT EDIT

package xpack

import (
	"context"
	"io"
	"strings"
)

func newCcrResumeFollowFunc(t Transport) CcrResumeFollow {
	return func(index string, o ...func(*CcrResumeFollowRequest)) (*Response, error) {
		var r = CcrResumeFollowRequest{Index: index}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/current/ccr-post-resume-follow.html.
//
type CcrResumeFollow func(index string, o ...func(*CcrResumeFollowRequest)) (*Response, error)

// CcrResumeFollowRequest configures the Ccr  Resume Follow API request.
//
type CcrResumeFollowRequest struct {
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
func (r CcrResumeFollowRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "POST"

	path.Grow(1 + len(r.Index) + 1 + len("_ccr") + 1 + len("resume_follow"))
	path.WriteString("/")
	path.WriteString(r.Index)
	path.WriteString("/")
	path.WriteString("_ccr")
	path.WriteString("/")
	path.WriteString("resume_follow")

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
func (f CcrResumeFollow) WithContext(v context.Context) func(*CcrResumeFollowRequest) {
	return func(r *CcrResumeFollowRequest) {
		r.ctx = v
	}
}

// WithBody - The name of the leader index and other optional ccr related parameters.
//
func (f CcrResumeFollow) WithBody(v io.Reader) func(*CcrResumeFollowRequest) {
	return func(r *CcrResumeFollowRequest) {
		r.Body = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f CcrResumeFollow) WithPretty() func(*CcrResumeFollowRequest) {
	return func(r *CcrResumeFollowRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f CcrResumeFollow) WithHuman() func(*CcrResumeFollowRequest) {
	return func(r *CcrResumeFollowRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f CcrResumeFollow) WithErrorTrace() func(*CcrResumeFollowRequest) {
	return func(r *CcrResumeFollowRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f CcrResumeFollow) WithFilterPath(v ...string) func(*CcrResumeFollowRequest) {
	return func(r *CcrResumeFollowRequest) {
		r.FilterPath = v
	}
}
