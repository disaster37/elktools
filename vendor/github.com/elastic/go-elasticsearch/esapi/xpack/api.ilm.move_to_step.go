// Code generated from specification version 7.0.0-SNAPSHOT: DO NOT EDIT

package xpack

import (
	"context"
	"io"
	"strings"
)

func newIlmMoveToStepFunc(t Transport) IlmMoveToStep {
	return func(o ...func(*IlmMoveToStepRequest)) (*Response, error) {
		var r = IlmMoveToStepRequest{}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/current/ilm-move-to-step.html.
//
type IlmMoveToStep func(o ...func(*IlmMoveToStepRequest)) (*Response, error)

// IlmMoveToStepRequest configures the Ilm   Move To Step API request.
//
type IlmMoveToStepRequest struct {
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
func (r IlmMoveToStepRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "POST"

	path.Grow(1 + len("_ilm") + 1 + len("move") + 1 + len(r.Index))
	path.WriteString("/")
	path.WriteString("_ilm")
	path.WriteString("/")
	path.WriteString("move")
	if r.Index != "" {
		path.WriteString("/")
		path.WriteString(r.Index)
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
func (f IlmMoveToStep) WithContext(v context.Context) func(*IlmMoveToStepRequest) {
	return func(r *IlmMoveToStepRequest) {
		r.ctx = v
	}
}

// WithIndex - the name of the index whose lifecycle step is to change.
//
func (f IlmMoveToStep) WithIndex(v string) func(*IlmMoveToStepRequest) {
	return func(r *IlmMoveToStepRequest) {
		r.Index = v
	}
}

// WithBody - The new lifecycle step to move to.
//
func (f IlmMoveToStep) WithBody(v io.Reader) func(*IlmMoveToStepRequest) {
	return func(r *IlmMoveToStepRequest) {
		r.Body = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f IlmMoveToStep) WithPretty() func(*IlmMoveToStepRequest) {
	return func(r *IlmMoveToStepRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f IlmMoveToStep) WithHuman() func(*IlmMoveToStepRequest) {
	return func(r *IlmMoveToStepRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f IlmMoveToStep) WithErrorTrace() func(*IlmMoveToStepRequest) {
	return func(r *IlmMoveToStepRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f IlmMoveToStep) WithFilterPath(v ...string) func(*IlmMoveToStepRequest) {
	return func(r *IlmMoveToStepRequest) {
		r.FilterPath = v
	}
}
