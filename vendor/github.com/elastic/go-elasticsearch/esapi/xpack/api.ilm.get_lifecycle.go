// Code generated from specification version 7.0.0-SNAPSHOT: DO NOT EDIT

package xpack

import (
	"context"
	"strings"
)

func newIlmGetLifecycleFunc(t Transport) IlmGetLifecycle {
	return func(o ...func(*IlmGetLifecycleRequest)) (*Response, error) {
		var r = IlmGetLifecycleRequest{}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/current/ilm-get-lifecycle.html.
//
type IlmGetLifecycle func(o ...func(*IlmGetLifecycleRequest)) (*Response, error)

// IlmGetLifecycleRequest configures the Ilm  Get Lifecycle API request.
//
type IlmGetLifecycleRequest struct {
	Policy string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r IlmGetLifecycleRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "GET"

	path.Grow(1 + len("_ilm") + 1 + len("policy") + 1 + len(r.Policy))
	path.WriteString("/")
	path.WriteString("_ilm")
	path.WriteString("/")
	path.WriteString("policy")
	if r.Policy != "" {
		path.WriteString("/")
		path.WriteString(r.Policy)
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
func (f IlmGetLifecycle) WithContext(v context.Context) func(*IlmGetLifecycleRequest) {
	return func(r *IlmGetLifecycleRequest) {
		r.ctx = v
	}
}

// WithPolicy - the name of the index lifecycle policy.
//
func (f IlmGetLifecycle) WithPolicy(v string) func(*IlmGetLifecycleRequest) {
	return func(r *IlmGetLifecycleRequest) {
		r.Policy = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f IlmGetLifecycle) WithPretty() func(*IlmGetLifecycleRequest) {
	return func(r *IlmGetLifecycleRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f IlmGetLifecycle) WithHuman() func(*IlmGetLifecycleRequest) {
	return func(r *IlmGetLifecycleRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f IlmGetLifecycle) WithErrorTrace() func(*IlmGetLifecycleRequest) {
	return func(r *IlmGetLifecycleRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f IlmGetLifecycle) WithFilterPath(v ...string) func(*IlmGetLifecycleRequest) {
	return func(r *IlmGetLifecycleRequest) {
		r.FilterPath = v
	}
}
