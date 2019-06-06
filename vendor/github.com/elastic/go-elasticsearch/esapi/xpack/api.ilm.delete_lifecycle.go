// Code generated from specification version 7.0.0-SNAPSHOT: DO NOT EDIT

package xpack

import (
	"context"
	"strings"
)

func newIlmDeleteLifecycleFunc(t Transport) IlmDeleteLifecycle {
	return func(o ...func(*IlmDeleteLifecycleRequest)) (*Response, error) {
		var r = IlmDeleteLifecycleRequest{}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/current/ilm-delete-lifecycle.html.
//
type IlmDeleteLifecycle func(o ...func(*IlmDeleteLifecycleRequest)) (*Response, error)

// IlmDeleteLifecycleRequest configures the Ilm  Delete Lifecycle API request.
//
type IlmDeleteLifecycleRequest struct {
	Policy string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r IlmDeleteLifecycleRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "DELETE"

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
func (f IlmDeleteLifecycle) WithContext(v context.Context) func(*IlmDeleteLifecycleRequest) {
	return func(r *IlmDeleteLifecycleRequest) {
		r.ctx = v
	}
}

// WithPolicy - the name of the index lifecycle policy.
//
func (f IlmDeleteLifecycle) WithPolicy(v string) func(*IlmDeleteLifecycleRequest) {
	return func(r *IlmDeleteLifecycleRequest) {
		r.Policy = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f IlmDeleteLifecycle) WithPretty() func(*IlmDeleteLifecycleRequest) {
	return func(r *IlmDeleteLifecycleRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f IlmDeleteLifecycle) WithHuman() func(*IlmDeleteLifecycleRequest) {
	return func(r *IlmDeleteLifecycleRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f IlmDeleteLifecycle) WithErrorTrace() func(*IlmDeleteLifecycleRequest) {
	return func(r *IlmDeleteLifecycleRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f IlmDeleteLifecycle) WithFilterPath(v ...string) func(*IlmDeleteLifecycleRequest) {
	return func(r *IlmDeleteLifecycleRequest) {
		r.FilterPath = v
	}
}
