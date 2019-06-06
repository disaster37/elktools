// Code generated from specification version 7.0.0-SNAPSHOT: DO NOT EDIT

package xpack

import (
	"context"
	"strings"
)

func newRollupGetRollupCapsFunc(t Transport) RollupGetRollupCaps {
	return func(o ...func(*RollupGetRollupCapsRequest)) (*Response, error) {
		var r = RollupGetRollupCapsRequest{}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

//
type RollupGetRollupCaps func(o ...func(*RollupGetRollupCapsRequest)) (*Response, error)

// RollupGetRollupCapsRequest configures the Rollup   Get Rollup Caps API request.
//
type RollupGetRollupCapsRequest struct {
	DocumentID string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r RollupGetRollupCapsRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "GET"

	path.Grow(1 + len("_rollup") + 1 + len("data") + 1 + len(r.DocumentID))
	path.WriteString("/")
	path.WriteString("_rollup")
	path.WriteString("/")
	path.WriteString("data")
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
func (f RollupGetRollupCaps) WithContext(v context.Context) func(*RollupGetRollupCapsRequest) {
	return func(r *RollupGetRollupCapsRequest) {
		r.ctx = v
	}
}

// WithDocumentID - the ID of the index to check rollup capabilities on, or left blank for all jobs.
//
func (f RollupGetRollupCaps) WithDocumentID(v string) func(*RollupGetRollupCapsRequest) {
	return func(r *RollupGetRollupCapsRequest) {
		r.DocumentID = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f RollupGetRollupCaps) WithPretty() func(*RollupGetRollupCapsRequest) {
	return func(r *RollupGetRollupCapsRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f RollupGetRollupCaps) WithHuman() func(*RollupGetRollupCapsRequest) {
	return func(r *RollupGetRollupCapsRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f RollupGetRollupCaps) WithErrorTrace() func(*RollupGetRollupCapsRequest) {
	return func(r *RollupGetRollupCapsRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f RollupGetRollupCaps) WithFilterPath(v ...string) func(*RollupGetRollupCapsRequest) {
	return func(r *RollupGetRollupCapsRequest) {
		r.FilterPath = v
	}
}
