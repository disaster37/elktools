// Code generated from specification version 7.0.0-SNAPSHOT: DO NOT EDIT

package xpack

import (
	"context"
	"io"
	"strings"
)

func newCcrPutAutoFollowPatternFunc(t Transport) CcrPutAutoFollowPattern {
	return func(body io.Reader, name string, o ...func(*CcrPutAutoFollowPatternRequest)) (*Response, error) {
		var r = CcrPutAutoFollowPatternRequest{Body: body, Name: name}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/current/ccr-put-auto-follow-pattern.html.
//
type CcrPutAutoFollowPattern func(body io.Reader, name string, o ...func(*CcrPutAutoFollowPatternRequest)) (*Response, error)

// CcrPutAutoFollowPatternRequest configures the Ccr    Put Auto Follow Pattern API request.
//
type CcrPutAutoFollowPatternRequest struct {
	Body io.Reader

	Name string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r CcrPutAutoFollowPatternRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "PUT"

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
func (f CcrPutAutoFollowPattern) WithContext(v context.Context) func(*CcrPutAutoFollowPatternRequest) {
	return func(r *CcrPutAutoFollowPatternRequest) {
		r.ctx = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f CcrPutAutoFollowPattern) WithPretty() func(*CcrPutAutoFollowPatternRequest) {
	return func(r *CcrPutAutoFollowPatternRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f CcrPutAutoFollowPattern) WithHuman() func(*CcrPutAutoFollowPatternRequest) {
	return func(r *CcrPutAutoFollowPatternRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f CcrPutAutoFollowPattern) WithErrorTrace() func(*CcrPutAutoFollowPatternRequest) {
	return func(r *CcrPutAutoFollowPatternRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f CcrPutAutoFollowPattern) WithFilterPath(v ...string) func(*CcrPutAutoFollowPatternRequest) {
	return func(r *CcrPutAutoFollowPatternRequest) {
		r.FilterPath = v
	}
}
