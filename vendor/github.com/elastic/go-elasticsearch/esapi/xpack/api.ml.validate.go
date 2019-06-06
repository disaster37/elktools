// Code generated from specification version 7.0.0-SNAPSHOT: DO NOT EDIT

package xpack

import (
	"context"
	"io"
	"strings"
)

func newMlValidateFunc(t Transport) MlValidate {
	return func(body io.Reader, o ...func(*MlValidateRequest)) (*Response, error) {
		var r = MlValidateRequest{Body: body}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

//
type MlValidate func(body io.Reader, o ...func(*MlValidateRequest)) (*Response, error)

// MlValidateRequest configures the Ml Validate API request.
//
type MlValidateRequest struct {
	Body io.Reader

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r MlValidateRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "POST"

	path.Grow(len("/_ml/anomaly_detectors/_validate"))
	path.WriteString("/_ml/anomaly_detectors/_validate")

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
func (f MlValidate) WithContext(v context.Context) func(*MlValidateRequest) {
	return func(r *MlValidateRequest) {
		r.ctx = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f MlValidate) WithPretty() func(*MlValidateRequest) {
	return func(r *MlValidateRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f MlValidate) WithHuman() func(*MlValidateRequest) {
	return func(r *MlValidateRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f MlValidate) WithErrorTrace() func(*MlValidateRequest) {
	return func(r *MlValidateRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f MlValidate) WithFilterPath(v ...string) func(*MlValidateRequest) {
	return func(r *MlValidateRequest) {
		r.FilterPath = v
	}
}
