// Code generated from specification version 7.0.0-SNAPSHOT: DO NOT EDIT

package xpack

import (
	"context"
	"io"
	"strings"
)

func newIlmPutLifecycleFunc(t Transport) IlmPutLifecycle {
	return func(o ...func(*IlmPutLifecycleRequest)) (*Response, error) {
		var r = IlmPutLifecycleRequest{}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/current/ilm-put-lifecycle.html.
//
type IlmPutLifecycle func(o ...func(*IlmPutLifecycleRequest)) (*Response, error)

// IlmPutLifecycleRequest configures the Ilm  Put Lifecycle API request.
//
type IlmPutLifecycleRequest struct {
	Body io.Reader

	Policy string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r IlmPutLifecycleRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "PUT"

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
func (f IlmPutLifecycle) WithContext(v context.Context) func(*IlmPutLifecycleRequest) {
	return func(r *IlmPutLifecycleRequest) {
		r.ctx = v
	}
}

// WithPolicy - the name of the index lifecycle policy.
//
func (f IlmPutLifecycle) WithPolicy(v string) func(*IlmPutLifecycleRequest) {
	return func(r *IlmPutLifecycleRequest) {
		r.Policy = v
	}
}

// WithBody - The lifecycle policy definition to register.
//
func (f IlmPutLifecycle) WithBody(v io.Reader) func(*IlmPutLifecycleRequest) {
	return func(r *IlmPutLifecycleRequest) {
		r.Body = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f IlmPutLifecycle) WithPretty() func(*IlmPutLifecycleRequest) {
	return func(r *IlmPutLifecycleRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f IlmPutLifecycle) WithHuman() func(*IlmPutLifecycleRequest) {
	return func(r *IlmPutLifecycleRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f IlmPutLifecycle) WithErrorTrace() func(*IlmPutLifecycleRequest) {
	return func(r *IlmPutLifecycleRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f IlmPutLifecycle) WithFilterPath(v ...string) func(*IlmPutLifecycleRequest) {
	return func(r *IlmPutLifecycleRequest) {
		r.FilterPath = v
	}
}
