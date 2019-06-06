// Code generated from specification version 7.0.0-SNAPSHOT: DO NOT EDIT

package xpack

import (
	"context"
	"strings"
)

func newCcrDeleteAutoFollowPatternFunc(t Transport) CcrDeleteAutoFollowPattern {
	return func(name string, o ...func(*CcrDeleteAutoFollowPatternRequest)) (*Response, error) {
		var r = CcrDeleteAutoFollowPatternRequest{Name: name}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/current/ccr-delete-auto-follow-pattern.html.
//
type CcrDeleteAutoFollowPattern func(name string, o ...func(*CcrDeleteAutoFollowPatternRequest)) (*Response, error)

// CcrDeleteAutoFollowPatternRequest configures the Ccr    Delete Auto Follow Pattern API request.
//
type CcrDeleteAutoFollowPatternRequest struct {
	Name string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r CcrDeleteAutoFollowPatternRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "DELETE"

	path.Grow(1 + len("_ccr") + 1 + len("auto_follow") + 1 + len(r.Name))
	path.WriteString("/")
	path.WriteString("_ccr")
	path.WriteString("/")
	path.WriteString("auto_follow")
	path.WriteString("/")
	path.WriteString(r.Name)

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
func (f CcrDeleteAutoFollowPattern) WithContext(v context.Context) func(*CcrDeleteAutoFollowPatternRequest) {
	return func(r *CcrDeleteAutoFollowPatternRequest) {
		r.ctx = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f CcrDeleteAutoFollowPattern) WithPretty() func(*CcrDeleteAutoFollowPatternRequest) {
	return func(r *CcrDeleteAutoFollowPatternRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f CcrDeleteAutoFollowPattern) WithHuman() func(*CcrDeleteAutoFollowPatternRequest) {
	return func(r *CcrDeleteAutoFollowPatternRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f CcrDeleteAutoFollowPattern) WithErrorTrace() func(*CcrDeleteAutoFollowPatternRequest) {
	return func(r *CcrDeleteAutoFollowPatternRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f CcrDeleteAutoFollowPattern) WithFilterPath(v ...string) func(*CcrDeleteAutoFollowPatternRequest) {
	return func(r *CcrDeleteAutoFollowPatternRequest) {
		r.FilterPath = v
	}
}
