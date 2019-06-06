// Code generated from specification version 7.0.0-SNAPSHOT: DO NOT EDIT

package xpack

import (
	"context"
	"strings"
	"time"
)

func newMlOpenJobFunc(t Transport) MlOpenJob {
	return func(job_id string, o ...func(*MlOpenJobRequest)) (*Response, error) {
		var r = MlOpenJobRequest{JobID: job_id}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

//
// See full documentation at http://www.elastic.co/guide/en/elasticsearch/reference/current/ml-open-job.html.
//
type MlOpenJob func(job_id string, o ...func(*MlOpenJobRequest)) (*Response, error)

// MlOpenJobRequest configures the Ml  Open Job API request.
//
type MlOpenJobRequest struct {
	JobID          string
	IgnoreDowntime *bool
	Timeout        time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r MlOpenJobRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "POST"

	path.Grow(1 + len("_ml") + 1 + len("anomaly_detectors") + 1 + len(r.JobID) + 1 + len("_open"))
	path.WriteString("/")
	path.WriteString("_ml")
	path.WriteString("/")
	path.WriteString("anomaly_detectors")
	path.WriteString("/")
	path.WriteString(r.JobID)
	path.WriteString("/")
	path.WriteString("_open")

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
func (f MlOpenJob) WithContext(v context.Context) func(*MlOpenJobRequest) {
	return func(r *MlOpenJobRequest) {
		r.ctx = v
	}
}

// WithIgnoreDowntime - controls if gaps in data are treated as anomalous or as a maintenance window after a job re-start.
//
func (f MlOpenJob) WithIgnoreDowntime(v bool) func(*MlOpenJobRequest) {
	return func(r *MlOpenJobRequest) {
		r.IgnoreDowntime = &v
	}
}

// WithTimeout - controls the time to wait until a job has opened. default to 30 minutes.
//
func (f MlOpenJob) WithTimeout(v time.Duration) func(*MlOpenJobRequest) {
	return func(r *MlOpenJobRequest) {
		r.Timeout = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f MlOpenJob) WithPretty() func(*MlOpenJobRequest) {
	return func(r *MlOpenJobRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f MlOpenJob) WithHuman() func(*MlOpenJobRequest) {
	return func(r *MlOpenJobRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f MlOpenJob) WithErrorTrace() func(*MlOpenJobRequest) {
	return func(r *MlOpenJobRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f MlOpenJob) WithFilterPath(v ...string) func(*MlOpenJobRequest) {
	return func(r *MlOpenJobRequest) {
		r.FilterPath = v
	}
}
