// Code generated from specification version 7.0.0-SNAPSHOT: DO NOT EDIT

package xpack

import (
	"context"
	"io"
	"strings"
)

func newMlPostDataFunc(t Transport) MlPostData {
	return func(body io.Reader, job_id string, o ...func(*MlPostDataRequest)) (*Response, error) {
		var r = MlPostDataRequest{Body: body, JobID: job_id}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

//
// See full documentation at http://www.elastic.co/guide/en/elasticsearch/reference/current/ml-post-data.html.
//
type MlPostData func(body io.Reader, job_id string, o ...func(*MlPostDataRequest)) (*Response, error)

// MlPostDataRequest configures the Ml  Post Data API request.
//
type MlPostDataRequest struct {
	Body io.Reader

	JobID      string
	ResetEnd   string
	ResetStart string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r MlPostDataRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "POST"

	path.Grow(1 + len("_ml") + 1 + len("anomaly_detectors") + 1 + len(r.JobID) + 1 + len("_data"))
	path.WriteString("/")
	path.WriteString("_ml")
	path.WriteString("/")
	path.WriteString("anomaly_detectors")
	path.WriteString("/")
	path.WriteString(r.JobID)
	path.WriteString("/")
	path.WriteString("_data")

	params = make(map[string]string)

	if r.ResetEnd != "" {
		params["reset_end"] = r.ResetEnd
	}

	if r.ResetStart != "" {
		params["reset_start"] = r.ResetStart
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
func (f MlPostData) WithContext(v context.Context) func(*MlPostDataRequest) {
	return func(r *MlPostDataRequest) {
		r.ctx = v
	}
}

// WithResetEnd - optional parameter to specify the end of the bucket resetting range.
//
func (f MlPostData) WithResetEnd(v string) func(*MlPostDataRequest) {
	return func(r *MlPostDataRequest) {
		r.ResetEnd = v
	}
}

// WithResetStart - optional parameter to specify the start of the bucket resetting range.
//
func (f MlPostData) WithResetStart(v string) func(*MlPostDataRequest) {
	return func(r *MlPostDataRequest) {
		r.ResetStart = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f MlPostData) WithPretty() func(*MlPostDataRequest) {
	return func(r *MlPostDataRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f MlPostData) WithHuman() func(*MlPostDataRequest) {
	return func(r *MlPostDataRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f MlPostData) WithErrorTrace() func(*MlPostDataRequest) {
	return func(r *MlPostDataRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f MlPostData) WithFilterPath(v ...string) func(*MlPostDataRequest) {
	return func(r *MlPostDataRequest) {
		r.FilterPath = v
	}
}
