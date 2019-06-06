// Code generated from specification version 7.0.0-SNAPSHOT: DO NOT EDIT

package xpack

import (
	"context"
	"strings"
)

func newRollupStartJobFunc(t Transport) RollupStartJob {
	return func(id string, o ...func(*RollupStartJobRequest)) (*Response, error) {
		var r = RollupStartJobRequest{DocumentID: id}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

//
type RollupStartJob func(id string, o ...func(*RollupStartJobRequest)) (*Response, error)

// RollupStartJobRequest configures the Rollup  Start Job API request.
//
type RollupStartJobRequest struct {
	DocumentID string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r RollupStartJobRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "POST"

	path.Grow(1 + len("_rollup") + 1 + len("job") + 1 + len(r.DocumentID) + 1 + len("_start"))
	path.WriteString("/")
	path.WriteString("_rollup")
	path.WriteString("/")
	path.WriteString("job")
	path.WriteString("/")
	path.WriteString(r.DocumentID)
	path.WriteString("/")
	path.WriteString("_start")

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
func (f RollupStartJob) WithContext(v context.Context) func(*RollupStartJobRequest) {
	return func(r *RollupStartJobRequest) {
		r.ctx = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f RollupStartJob) WithPretty() func(*RollupStartJobRequest) {
	return func(r *RollupStartJobRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f RollupStartJob) WithHuman() func(*RollupStartJobRequest) {
	return func(r *RollupStartJobRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f RollupStartJob) WithErrorTrace() func(*RollupStartJobRequest) {
	return func(r *RollupStartJobRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f RollupStartJob) WithFilterPath(v ...string) func(*RollupStartJobRequest) {
	return func(r *RollupStartJobRequest) {
		r.FilterPath = v
	}
}
