// Code generated from specification version 7.0.0-SNAPSHOT: DO NOT EDIT

package xpack

import (
	"context"
	"strings"
)

func newCcrGetAutoFollowPatternFunc(t Transport) CcrGetAutoFollowPattern {
	return func(o ...func(*CcrGetAutoFollowPatternRequest)) (*Response, error) {
		var r = CcrGetAutoFollowPatternRequest{}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/current/ccr-get-auto-follow-pattern.html.
//
type CcrGetAutoFollowPattern func(o ...func(*CcrGetAutoFollowPatternRequest)) (*Response, error)

// CcrGetAutoFollowPatternRequest configures the Ccr    Get Auto Follow Pattern API request.
//
type CcrGetAutoFollowPatternRequest struct {
	Name string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r CcrGetAutoFollowPatternRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "GET"

	path.Grow(1 + len("_ccr") + 1 + len("auto_follow") + 1 + len(r.Name))
	path.WriteString("/")
	path.WriteString("_ccr")
	path.WriteString("/")
	path.WriteString("auto_follow")
	if r.Name != "" {
		path.WriteString("/")
		path.WriteString(r.Name)
	}

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
func (f CcrGetAutoFollowPattern) WithContext(v context.Context) func(*CcrGetAutoFollowPatternRequest) {
	return func(r *CcrGetAutoFollowPatternRequest) {
		r.ctx = v
	}
}

// WithName - the name of the auto follow pattern..
//
func (f CcrGetAutoFollowPattern) WithName(v string) func(*CcrGetAutoFollowPatternRequest) {
	return func(r *CcrGetAutoFollowPatternRequest) {
		r.Name = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f CcrGetAutoFollowPattern) WithPretty() func(*CcrGetAutoFollowPatternRequest) {
	return func(r *CcrGetAutoFollowPatternRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f CcrGetAutoFollowPattern) WithHuman() func(*CcrGetAutoFollowPatternRequest) {
	return func(r *CcrGetAutoFollowPatternRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f CcrGetAutoFollowPattern) WithErrorTrace() func(*CcrGetAutoFollowPatternRequest) {
	return func(r *CcrGetAutoFollowPatternRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f CcrGetAutoFollowPattern) WithFilterPath(v ...string) func(*CcrGetAutoFollowPatternRequest) {
	return func(r *CcrGetAutoFollowPatternRequest) {
		r.FilterPath = v
	}
}
