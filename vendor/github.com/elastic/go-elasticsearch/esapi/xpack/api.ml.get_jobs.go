// Code generated from specification version 7.0.0-SNAPSHOT: DO NOT EDIT

package xpack

import (
	"context"
	"strconv"
	"strings"
)

func newMlGetJobsFunc(t Transport) MlGetJobs {
	return func(o ...func(*MlGetJobsRequest)) (*Response, error) {
		var r = MlGetJobsRequest{}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

//
// See full documentation at http://www.elastic.co/guide/en/elasticsearch/reference/current/ml-get-job.html.
//
type MlGetJobs func(o ...func(*MlGetJobsRequest)) (*Response, error)

// MlGetJobsRequest configures the Ml  Get Jobs API request.
//
type MlGetJobsRequest struct {
	JobID       string
	AllowNoJobs *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r MlGetJobsRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "GET"

	path.Grow(1 + len("_ml") + 1 + len("anomaly_detectors") + 1 + len(r.JobID))
	path.WriteString("/")
	path.WriteString("_ml")
	path.WriteString("/")
	path.WriteString("anomaly_detectors")
	if r.JobID != "" {
		path.WriteString("/")
		path.WriteString(r.JobID)
	}

	params = make(map[string]string)

	if r.AllowNoJobs != nil {
		params["allow_no_jobs"] = strconv.FormatBool(*r.AllowNoJobs)
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
func (f MlGetJobs) WithContext(v context.Context) func(*MlGetJobsRequest) {
	return func(r *MlGetJobsRequest) {
		r.ctx = v
	}
}

// WithJobID - the ID of the jobs to fetch.
//
func (f MlGetJobs) WithJobID(v string) func(*MlGetJobsRequest) {
	return func(r *MlGetJobsRequest) {
		r.JobID = v
	}
}

// WithAllowNoJobs - whether to ignore if a wildcard expression matches no jobs. (this includes `_all` string or when no jobs have been specified).
//
func (f MlGetJobs) WithAllowNoJobs(v bool) func(*MlGetJobsRequest) {
	return func(r *MlGetJobsRequest) {
		r.AllowNoJobs = &v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f MlGetJobs) WithPretty() func(*MlGetJobsRequest) {
	return func(r *MlGetJobsRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f MlGetJobs) WithHuman() func(*MlGetJobsRequest) {
	return func(r *MlGetJobsRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f MlGetJobs) WithErrorTrace() func(*MlGetJobsRequest) {
	return func(r *MlGetJobsRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f MlGetJobs) WithFilterPath(v ...string) func(*MlGetJobsRequest) {
	return func(r *MlGetJobsRequest) {
		r.FilterPath = v
	}
}
