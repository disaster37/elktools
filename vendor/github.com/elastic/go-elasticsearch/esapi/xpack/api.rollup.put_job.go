// Code generated from specification version 7.0.0-SNAPSHOT: DO NOT EDIT

package xpack

import (
	"context"
	"io"
	"strings"
)

func newRollupPutJobFunc(t Transport) RollupPutJob {
	return func(id string, body io.Reader, o ...func(*RollupPutJobRequest)) (*Response, error) {
		var r = RollupPutJobRequest{DocumentID: id, Body: body}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

//
type RollupPutJob func(id string, body io.Reader, o ...func(*RollupPutJobRequest)) (*Response, error)

// RollupPutJobRequest configures the Rollup  Put Job API request.
//
type RollupPutJobRequest struct {
	DocumentID string
	Body       io.Reader

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r RollupPutJobRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "PUT"

	path.Grow(1 + len("_rollup") + 1 + len("job") + 1 + len(r.DocumentID))
	path.WriteString("/")
	path.WriteString("_rollup")
	path.WriteString("/")
	path.WriteString("job")
	path.WriteString("/")
	path.WriteString(r.DocumentID)

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
func (f RollupPutJob) WithContext(v context.Context) func(*RollupPutJobRequest) {
	return func(r *RollupPutJobRequest) {
		r.ctx = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f RollupPutJob) WithPretty() func(*RollupPutJobRequest) {
	return func(r *RollupPutJobRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f RollupPutJob) WithHuman() func(*RollupPutJobRequest) {
	return func(r *RollupPutJobRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f RollupPutJob) WithErrorTrace() func(*RollupPutJobRequest) {
	return func(r *RollupPutJobRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f RollupPutJob) WithFilterPath(v ...string) func(*RollupPutJobRequest) {
	return func(r *RollupPutJobRequest) {
		r.FilterPath = v
	}
}
