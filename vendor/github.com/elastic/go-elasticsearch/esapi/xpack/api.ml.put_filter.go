// Code generated from specification version 7.0.0-SNAPSHOT: DO NOT EDIT

package xpack

import (
	"context"
	"io"
	"strings"
)

func newMlPutFilterFunc(t Transport) MlPutFilter {
	return func(body io.Reader, filter_id string, o ...func(*MlPutFilterRequest)) (*Response, error) {
		var r = MlPutFilterRequest{Body: body, FilterID: filter_id}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

//
type MlPutFilter func(body io.Reader, filter_id string, o ...func(*MlPutFilterRequest)) (*Response, error)

// MlPutFilterRequest configures the Ml  Put Filter API request.
//
type MlPutFilterRequest struct {
	Body io.Reader

	FilterID string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r MlPutFilterRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "PUT"

	path.Grow(1 + len("_ml") + 1 + len("filters") + 1 + len(r.FilterID))
	path.WriteString("/")
	path.WriteString("_ml")
	path.WriteString("/")
	path.WriteString("filters")
	path.WriteString("/")
	path.WriteString(r.FilterID)

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
func (f MlPutFilter) WithContext(v context.Context) func(*MlPutFilterRequest) {
	return func(r *MlPutFilterRequest) {
		r.ctx = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f MlPutFilter) WithPretty() func(*MlPutFilterRequest) {
	return func(r *MlPutFilterRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f MlPutFilter) WithHuman() func(*MlPutFilterRequest) {
	return func(r *MlPutFilterRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f MlPutFilter) WithErrorTrace() func(*MlPutFilterRequest) {
	return func(r *MlPutFilterRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f MlPutFilter) WithFilterPath(v ...string) func(*MlPutFilterRequest) {
	return func(r *MlPutFilterRequest) {
		r.FilterPath = v
	}
}
