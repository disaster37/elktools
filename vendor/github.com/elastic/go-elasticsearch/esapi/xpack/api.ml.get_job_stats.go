// Code generated from specification version 7.0.0-SNAPSHOT: DO NOT EDIT

package xpack

import (
	"context"
	"strconv"
	"strings"
)

func newMlGetJobStatsFunc(t Transport) MlGetJobStats {
	return func(o ...func(*MlGetJobStatsRequest)) (*Response, error) {
		var r = MlGetJobStatsRequest{}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

//
// See full documentation at http://www.elastic.co/guide/en/elasticsearch/reference/current/ml-get-job-stats.html.
//
type MlGetJobStats func(o ...func(*MlGetJobStatsRequest)) (*Response, error)

// MlGetJobStatsRequest configures the Ml   Get Job Stats API request.
//
type MlGetJobStatsRequest struct {
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
func (r MlGetJobStatsRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "GET"

	path.Grow(1 + len("_ml") + 1 + len("anomaly_detectors") + 1 + len(r.JobID) + 1 + len("_stats"))
	path.WriteString("/")
	path.WriteString("_ml")
	path.WriteString("/")
	path.WriteString("anomaly_detectors")
	if r.JobID != "" {
		path.WriteString("/")
		path.WriteString(r.JobID)
	}
	path.WriteString("/")
	path.WriteString("_stats")

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
func (f MlGetJobStats) WithContext(v context.Context) func(*MlGetJobStatsRequest) {
	return func(r *MlGetJobStatsRequest) {
		r.ctx = v
	}
}

// WithJobID - the ID of the jobs stats to fetch.
//
func (f MlGetJobStats) WithJobID(v string) func(*MlGetJobStatsRequest) {
	return func(r *MlGetJobStatsRequest) {
		r.JobID = v
	}
}

// WithAllowNoJobs - whether to ignore if a wildcard expression matches no jobs. (this includes `_all` string or when no jobs have been specified).
//
func (f MlGetJobStats) WithAllowNoJobs(v bool) func(*MlGetJobStatsRequest) {
	return func(r *MlGetJobStatsRequest) {
		r.AllowNoJobs = &v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f MlGetJobStats) WithPretty() func(*MlGetJobStatsRequest) {
	return func(r *MlGetJobStatsRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f MlGetJobStats) WithHuman() func(*MlGetJobStatsRequest) {
	return func(r *MlGetJobStatsRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f MlGetJobStats) WithErrorTrace() func(*MlGetJobStatsRequest) {
	return func(r *MlGetJobStatsRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f MlGetJobStats) WithFilterPath(v ...string) func(*MlGetJobStatsRequest) {
	return func(r *MlGetJobStatsRequest) {
		r.FilterPath = v
	}
}
