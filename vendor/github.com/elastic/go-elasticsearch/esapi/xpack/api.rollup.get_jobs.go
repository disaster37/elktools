// Code generated from specification version 7.0.0-SNAPSHOT: DO NOT EDIT

package xpack

import (
	"context"
	"strings"
)

func newRollupGetJobsFunc(t Transport) RollupGetJobs {
	return func(o ...func(*RollupGetJobsRequest)) (*Response, error) {
		var r = RollupGetJobsRequest{}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

//
type RollupGetJobs func(o ...func(*RollupGetJobsRequest)) (*Response, error)

// RollupGetJobsRequest configures the Rollup  Get Jobs API request.
//
type RollupGetJobsRequest struct {
	DocumentID string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r RollupGetJobsRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "GET"

	path.Grow(1 + len("_rollup") + 1 + len("job") + 1 + len(r.DocumentID))
	path.WriteString("/")
	path.WriteString("_rollup")
	path.WriteString("/")
	path.WriteString("job")
	if r.DocumentID != "" {
		path.WriteString("/")
		path.WriteString(r.DocumentID)
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
func (f RollupGetJobs) WithContext(v context.Context) func(*RollupGetJobsRequest) {
	return func(r *RollupGetJobsRequest) {
		r.ctx = v
	}
}

// WithDocumentID - the ID of the job(s) to fetch. accepts glob patterns, or left blank for all jobs.
//
func (f RollupGetJobs) WithDocumentID(v string) func(*RollupGetJobsRequest) {
	return func(r *RollupGetJobsRequest) {
		r.DocumentID = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f RollupGetJobs) WithPretty() func(*RollupGetJobsRequest) {
	return func(r *RollupGetJobsRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f RollupGetJobs) WithHuman() func(*RollupGetJobsRequest) {
	return func(r *RollupGetJobsRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f RollupGetJobs) WithErrorTrace() func(*RollupGetJobsRequest) {
	return func(r *RollupGetJobsRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f RollupGetJobs) WithFilterPath(v ...string) func(*RollupGetJobsRequest) {
	return func(r *RollupGetJobsRequest) {
		r.FilterPath = v
	}
}
