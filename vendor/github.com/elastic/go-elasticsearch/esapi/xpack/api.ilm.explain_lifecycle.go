// Code generated from specification version 7.0.0-SNAPSHOT: DO NOT EDIT

package xpack

import (
	"context"
	"strings"
)

func newIlmExplainLifecycleFunc(t Transport) IlmExplainLifecycle {
	return func(o ...func(*IlmExplainLifecycleRequest)) (*Response, error) {
		var r = IlmExplainLifecycleRequest{}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/current/ilm-explain-lifecycle.html.
//
type IlmExplainLifecycle func(o ...func(*IlmExplainLifecycleRequest)) (*Response, error)

// IlmExplainLifecycleRequest configures the Ilm  Explain Lifecycle API request.
//
type IlmExplainLifecycleRequest struct {
	Index string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r IlmExplainLifecycleRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "GET"

	path.Grow(1 + len(r.Index) + 1 + len("_ilm") + 1 + len("explain"))
	if r.Index != "" {
		path.WriteString("/")
		path.WriteString(r.Index)
	}
	path.WriteString("/")
	path.WriteString("_ilm")
	path.WriteString("/")
	path.WriteString("explain")

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
func (f IlmExplainLifecycle) WithContext(v context.Context) func(*IlmExplainLifecycleRequest) {
	return func(r *IlmExplainLifecycleRequest) {
		r.ctx = v
	}
}

// WithIndex - the name of the index to explain.
//
func (f IlmExplainLifecycle) WithIndex(v string) func(*IlmExplainLifecycleRequest) {
	return func(r *IlmExplainLifecycleRequest) {
		r.Index = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f IlmExplainLifecycle) WithPretty() func(*IlmExplainLifecycleRequest) {
	return func(r *IlmExplainLifecycleRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f IlmExplainLifecycle) WithHuman() func(*IlmExplainLifecycleRequest) {
	return func(r *IlmExplainLifecycleRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f IlmExplainLifecycle) WithErrorTrace() func(*IlmExplainLifecycleRequest) {
	return func(r *IlmExplainLifecycleRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f IlmExplainLifecycle) WithFilterPath(v ...string) func(*IlmExplainLifecycleRequest) {
	return func(r *IlmExplainLifecycleRequest) {
		r.FilterPath = v
	}
}
