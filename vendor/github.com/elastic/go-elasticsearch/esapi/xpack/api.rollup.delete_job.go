// Code generated from specification version 7.0.0-SNAPSHOT: DO NOT EDIT

package xpack

import (
	"context"
	"strings"
)

func newRollupDeleteJobFunc(t Transport) RollupDeleteJob {
	return func(id string, o ...func(*RollupDeleteJobRequest)) (*Response, error) {
		var r = RollupDeleteJobRequest{DocumentID: id}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

//
type RollupDeleteJob func(id string, o ...func(*RollupDeleteJobRequest)) (*Response, error)

// RollupDeleteJobRequest configures the Rollup  Delete Job API request.
//
type RollupDeleteJobRequest struct {
	DocumentID string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r RollupDeleteJobRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "DELETE"

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
func (f RollupDeleteJob) WithContext(v context.Context) func(*RollupDeleteJobRequest) {
	return func(r *RollupDeleteJobRequest) {
		r.ctx = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f RollupDeleteJob) WithPretty() func(*RollupDeleteJobRequest) {
	return func(r *RollupDeleteJobRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f RollupDeleteJob) WithHuman() func(*RollupDeleteJobRequest) {
	return func(r *RollupDeleteJobRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f RollupDeleteJob) WithErrorTrace() func(*RollupDeleteJobRequest) {
	return func(r *RollupDeleteJobRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f RollupDeleteJob) WithFilterPath(v ...string) func(*RollupDeleteJobRequest) {
	return func(r *RollupDeleteJobRequest) {
		r.FilterPath = v
	}
}
