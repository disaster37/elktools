// Code generated from specification version 7.0.0-SNAPSHOT: DO NOT EDIT

package xpack

import (
	"context"
	"io"
	"strconv"
	"strings"
	"time"
)

func newMlCloseJobFunc(t Transport) MlCloseJob {
	return func(job_id string, o ...func(*MlCloseJobRequest)) (*Response, error) {
		var r = MlCloseJobRequest{JobID: job_id}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

//
// See full documentation at http://www.elastic.co/guide/en/elasticsearch/reference/current/ml-close-job.html.
//
type MlCloseJob func(job_id string, o ...func(*MlCloseJobRequest)) (*Response, error)

// MlCloseJobRequest configures the Ml  Close Job API request.
//
type MlCloseJobRequest struct {
	Body io.Reader

	JobID       string
	AllowNoJobs *bool
	Force       *bool
	Timeout     time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r MlCloseJobRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "POST"

	path.Grow(1 + len("_ml") + 1 + len("anomaly_detectors") + 1 + len(r.JobID) + 1 + len("_close"))
	path.WriteString("/")
	path.WriteString("_ml")
	path.WriteString("/")
	path.WriteString("anomaly_detectors")
	path.WriteString("/")
	path.WriteString(r.JobID)
	path.WriteString("/")
	path.WriteString("_close")

	params = make(map[string]string)

	if r.AllowNoJobs != nil {
		params["allow_no_jobs"] = strconv.FormatBool(*r.AllowNoJobs)
	}

	if r.Force != nil {
		params["force"] = strconv.FormatBool(*r.Force)
	}

	if r.Timeout != 0 {
		params["timeout"] = formatDuration(r.Timeout)
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
func (f MlCloseJob) WithContext(v context.Context) func(*MlCloseJobRequest) {
	return func(r *MlCloseJobRequest) {
		r.ctx = v
	}
}

// WithBody - The URL params optionally sent in the body.
//
func (f MlCloseJob) WithBody(v io.Reader) func(*MlCloseJobRequest) {
	return func(r *MlCloseJobRequest) {
		r.Body = v
	}
}

// WithAllowNoJobs - whether to ignore if a wildcard expression matches no jobs. (this includes `_all` string or when no jobs have been specified).
//
func (f MlCloseJob) WithAllowNoJobs(v bool) func(*MlCloseJobRequest) {
	return func(r *MlCloseJobRequest) {
		r.AllowNoJobs = &v
	}
}

// WithForce - true if the job should be forcefully closed.
//
func (f MlCloseJob) WithForce(v bool) func(*MlCloseJobRequest) {
	return func(r *MlCloseJobRequest) {
		r.Force = &v
	}
}

// WithTimeout - controls the time to wait until a job has closed. default to 30 minutes.
//
func (f MlCloseJob) WithTimeout(v time.Duration) func(*MlCloseJobRequest) {
	return func(r *MlCloseJobRequest) {
		r.Timeout = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f MlCloseJob) WithPretty() func(*MlCloseJobRequest) {
	return func(r *MlCloseJobRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f MlCloseJob) WithHuman() func(*MlCloseJobRequest) {
	return func(r *MlCloseJobRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f MlCloseJob) WithErrorTrace() func(*MlCloseJobRequest) {
	return func(r *MlCloseJobRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f MlCloseJob) WithFilterPath(v ...string) func(*MlCloseJobRequest) {
	return func(r *MlCloseJobRequest) {
		r.FilterPath = v
	}
}
