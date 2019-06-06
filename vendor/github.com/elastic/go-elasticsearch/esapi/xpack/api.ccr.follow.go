// Code generated from specification version 7.0.0-SNAPSHOT: DO NOT EDIT

package xpack

import (
	"context"
	"io"
	"strings"
)

func newCcrFollowFunc(t Transport) CcrFollow {
	return func(index string, body io.Reader, o ...func(*CcrFollowRequest)) (*Response, error) {
		var r = CcrFollowRequest{Index: index, Body: body}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/current/ccr-put-follow.html.
//
type CcrFollow func(index string, body io.Reader, o ...func(*CcrFollowRequest)) (*Response, error)

// CcrFollowRequest configures the Ccr Follow API request.
//
type CcrFollowRequest struct {
	Index string
	Body  io.Reader

	WaitForActiveShards string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r CcrFollowRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "PUT"

	path.Grow(1 + len(r.Index) + 1 + len("_ccr") + 1 + len("follow"))
	path.WriteString("/")
	path.WriteString(r.Index)
	path.WriteString("/")
	path.WriteString("_ccr")
	path.WriteString("/")
	path.WriteString("follow")

	params = make(map[string]string)

	if r.WaitForActiveShards != "" {
		params["wait_for_active_shards"] = r.WaitForActiveShards
	}

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
func (f CcrFollow) WithContext(v context.Context) func(*CcrFollowRequest) {
	return func(r *CcrFollowRequest) {
		r.ctx = v
	}
}

// WithWaitForActiveShards - sets the number of shard copies that must be active before returning. defaults to 0. set to `all` for all shard copies, otherwise set to any non-negative value less than or equal to the total number of copies for the shard (number of replicas + 1).
//
func (f CcrFollow) WithWaitForActiveShards(v string) func(*CcrFollowRequest) {
	return func(r *CcrFollowRequest) {
		r.WaitForActiveShards = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f CcrFollow) WithPretty() func(*CcrFollowRequest) {
	return func(r *CcrFollowRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f CcrFollow) WithHuman() func(*CcrFollowRequest) {
	return func(r *CcrFollowRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f CcrFollow) WithErrorTrace() func(*CcrFollowRequest) {
	return func(r *CcrFollowRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f CcrFollow) WithFilterPath(v ...string) func(*CcrFollowRequest) {
	return func(r *CcrFollowRequest) {
		r.FilterPath = v
	}
}
