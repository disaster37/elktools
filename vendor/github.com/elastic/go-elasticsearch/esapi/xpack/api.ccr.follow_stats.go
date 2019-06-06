// Code generated from specification version 7.0.0-SNAPSHOT: DO NOT EDIT

package xpack

import (
	"context"
	"strings"
)

func newCcrFollowStatsFunc(t Transport) CcrFollowStats {
	return func(index []string, o ...func(*CcrFollowStatsRequest)) (*Response, error) {
		var r = CcrFollowStatsRequest{Index: index}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/current/ccr-get-follow-stats.html.
//
type CcrFollowStats func(index []string, o ...func(*CcrFollowStatsRequest)) (*Response, error)

// CcrFollowStatsRequest configures the Ccr  Follow Stats API request.
//
type CcrFollowStatsRequest struct {
	Index []string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r CcrFollowStatsRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "GET"

	path.Grow(1 + len(strings.Join(r.Index, ",")) + 1 + len("_ccr") + 1 + len("stats"))
	path.WriteString("/")
	path.WriteString(strings.Join(r.Index, ","))
	path.WriteString("/")
	path.WriteString("_ccr")
	path.WriteString("/")
	path.WriteString("stats")

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
func (f CcrFollowStats) WithContext(v context.Context) func(*CcrFollowStatsRequest) {
	return func(r *CcrFollowStatsRequest) {
		r.ctx = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f CcrFollowStats) WithPretty() func(*CcrFollowStatsRequest) {
	return func(r *CcrFollowStatsRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f CcrFollowStats) WithHuman() func(*CcrFollowStatsRequest) {
	return func(r *CcrFollowStatsRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f CcrFollowStats) WithErrorTrace() func(*CcrFollowStatsRequest) {
	return func(r *CcrFollowStatsRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f CcrFollowStats) WithFilterPath(v ...string) func(*CcrFollowStatsRequest) {
	return func(r *CcrFollowStatsRequest) {
		r.FilterPath = v
	}
}
