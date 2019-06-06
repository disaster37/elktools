// Code generated from specification version 7.0.0-SNAPSHOT: DO NOT EDIT

package xpack

import (
	"context"
	"io"
	"strings"
)

func newMlUpdateJobFunc(t Transport) MlUpdateJob {
	return func(body io.Reader, job_id string, o ...func(*MlUpdateJobRequest)) (*Response, error) {
		var r = MlUpdateJobRequest{Body: body, JobID: job_id}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

//
// See full documentation at http://www.elastic.co/guide/en/elasticsearch/reference/current/ml-update-job.html.
//
type MlUpdateJob func(body io.Reader, job_id string, o ...func(*MlUpdateJobRequest)) (*Response, error)

// MlUpdateJobRequest configures the Ml  Update Job API request.
//
type MlUpdateJobRequest struct {
	Body io.Reader

	JobID string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r MlUpdateJobRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "POST"

	path.Grow(1 + len("_ml") + 1 + len("anomaly_detectors") + 1 + len(r.JobID) + 1 + len("_update"))
	path.WriteString("/")
	path.WriteString("_ml")
	path.WriteString("/")
	path.WriteString("anomaly_detectors")
	path.WriteString("/")
	path.WriteString(r.JobID)
	path.WriteString("/")
	path.WriteString("_update")

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
func (f MlUpdateJob) WithContext(v context.Context) func(*MlUpdateJobRequest) {
	return func(r *MlUpdateJobRequest) {
		r.ctx = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f MlUpdateJob) WithPretty() func(*MlUpdateJobRequest) {
	return func(r *MlUpdateJobRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f MlUpdateJob) WithHuman() func(*MlUpdateJobRequest) {
	return func(r *MlUpdateJobRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f MlUpdateJob) WithErrorTrace() func(*MlUpdateJobRequest) {
	return func(r *MlUpdateJobRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f MlUpdateJob) WithFilterPath(v ...string) func(*MlUpdateJobRequest) {
	return func(r *MlUpdateJobRequest) {
		r.FilterPath = v
	}
}
