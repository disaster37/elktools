// Code generated from specification version 7.0.0-SNAPSHOT: DO NOT EDIT

package xpack

import (
	"context"
	"io"
	"strings"
)

func newMlValidateDetectorFunc(t Transport) MlValidateDetector {
	return func(body io.Reader, o ...func(*MlValidateDetectorRequest)) (*Response, error) {
		var r = MlValidateDetectorRequest{Body: body}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

//
type MlValidateDetector func(body io.Reader, o ...func(*MlValidateDetectorRequest)) (*Response, error)

// MlValidateDetectorRequest configures the Ml  Validate Detector API request.
//
type MlValidateDetectorRequest struct {
	Body io.Reader

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r MlValidateDetectorRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "POST"

	path.Grow(len("/_ml/anomaly_detectors/_validate/detector"))
	path.WriteString("/_ml/anomaly_detectors/_validate/detector")

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
func (f MlValidateDetector) WithContext(v context.Context) func(*MlValidateDetectorRequest) {
	return func(r *MlValidateDetectorRequest) {
		r.ctx = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f MlValidateDetector) WithPretty() func(*MlValidateDetectorRequest) {
	return func(r *MlValidateDetectorRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f MlValidateDetector) WithHuman() func(*MlValidateDetectorRequest) {
	return func(r *MlValidateDetectorRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f MlValidateDetector) WithErrorTrace() func(*MlValidateDetectorRequest) {
	return func(r *MlValidateDetectorRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f MlValidateDetector) WithFilterPath(v ...string) func(*MlValidateDetectorRequest) {
	return func(r *MlValidateDetectorRequest) {
		r.FilterPath = v
	}
}
