// Code generated from specification version 7.0.0-SNAPSHOT: DO NOT EDIT

package xpack

import (
	"context"
	"strings"
)

func newCcrPauseFollowFunc(t Transport) CcrPauseFollow {
	return func(index string, o ...func(*CcrPauseFollowRequest)) (*Response, error) {
		var r = CcrPauseFollowRequest{Index: index}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/current/ccr-post-pause-follow.html.
//
type CcrPauseFollow func(index string, o ...func(*CcrPauseFollowRequest)) (*Response, error)

// CcrPauseFollowRequest configures the Ccr  Pause Follow API request.
//
type CcrPauseFollowRequest struct {
	Index string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r CcrPauseFollowRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "POST"

	path.Grow(1 + len(r.Index) + 1 + len("_ccr") + 1 + len("pause_follow"))
	path.WriteString("/")
	path.WriteString(r.Index)
	path.WriteString("/")
	path.WriteString("_ccr")
	path.WriteString("/")
	path.WriteString("pause_follow")

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
func (f CcrPauseFollow) WithContext(v context.Context) func(*CcrPauseFollowRequest) {
	return func(r *CcrPauseFollowRequest) {
		r.ctx = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f CcrPauseFollow) WithPretty() func(*CcrPauseFollowRequest) {
	return func(r *CcrPauseFollowRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f CcrPauseFollow) WithHuman() func(*CcrPauseFollowRequest) {
	return func(r *CcrPauseFollowRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f CcrPauseFollow) WithErrorTrace() func(*CcrPauseFollowRequest) {
	return func(r *CcrPauseFollowRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f CcrPauseFollow) WithFilterPath(v ...string) func(*CcrPauseFollowRequest) {
	return func(r *CcrPauseFollowRequest) {
		r.FilterPath = v
	}
}
