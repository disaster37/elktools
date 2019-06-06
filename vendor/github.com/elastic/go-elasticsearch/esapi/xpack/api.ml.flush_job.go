// Code generated from specification version 7.0.0-SNAPSHOT: DO NOT EDIT

package xpack

import (
	"context"
	"io"
	"strconv"
	"strings"
)

func newMlFlushJobFunc(t Transport) MlFlushJob {
	return func(job_id string, o ...func(*MlFlushJobRequest)) (*Response, error) {
		var r = MlFlushJobRequest{JobID: job_id}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

//
// See full documentation at http://www.elastic.co/guide/en/elasticsearch/reference/current/ml-flush-job.html.
//
type MlFlushJob func(job_id string, o ...func(*MlFlushJobRequest)) (*Response, error)

// MlFlushJobRequest configures the Ml  Flush Job API request.
//
type MlFlushJobRequest struct {
	Body io.Reader

	JobID       string
	AdvanceTime string
	CalcInterim *bool
	End         string
	SkipTime    string
	Start       string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r MlFlushJobRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "POST"

	path.Grow(1 + len("_ml") + 1 + len("anomaly_detectors") + 1 + len(r.JobID) + 1 + len("_flush"))
	path.WriteString("/")
	path.WriteString("_ml")
	path.WriteString("/")
	path.WriteString("anomaly_detectors")
	path.WriteString("/")
	path.WriteString(r.JobID)
	path.WriteString("/")
	path.WriteString("_flush")

	params = make(map[string]string)

	if r.AdvanceTime != "" {
		params["advance_time"] = r.AdvanceTime
	}

	if r.CalcInterim != nil {
		params["calc_interim"] = strconv.FormatBool(*r.CalcInterim)
	}

	if r.End != "" {
		params["end"] = r.End
	}

	if r.SkipTime != "" {
		params["skip_time"] = r.SkipTime
	}

	if r.Start != "" {
		params["start"] = r.Start
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
func (f MlFlushJob) WithContext(v context.Context) func(*MlFlushJobRequest) {
	return func(r *MlFlushJobRequest) {
		r.ctx = v
	}
}

// WithBody - Flush parameters.
//
func (f MlFlushJob) WithBody(v io.Reader) func(*MlFlushJobRequest) {
	return func(r *MlFlushJobRequest) {
		r.Body = v
	}
}

// WithAdvanceTime - advances time to the given value generating results and updating the model for the advanced interval.
//
func (f MlFlushJob) WithAdvanceTime(v string) func(*MlFlushJobRequest) {
	return func(r *MlFlushJobRequest) {
		r.AdvanceTime = v
	}
}

// WithCalcInterim - calculates interim results for the most recent bucket or all buckets within the latency period.
//
func (f MlFlushJob) WithCalcInterim(v bool) func(*MlFlushJobRequest) {
	return func(r *MlFlushJobRequest) {
		r.CalcInterim = &v
	}
}

// WithEnd - when used in conjunction with calc_interim, specifies the range of buckets on which to calculate interim results.
//
func (f MlFlushJob) WithEnd(v string) func(*MlFlushJobRequest) {
	return func(r *MlFlushJobRequest) {
		r.End = v
	}
}

// WithSkipTime - skips time to the given value without generating results or updating the model for the skipped interval.
//
func (f MlFlushJob) WithSkipTime(v string) func(*MlFlushJobRequest) {
	return func(r *MlFlushJobRequest) {
		r.SkipTime = v
	}
}

// WithStart - when used in conjunction with calc_interim, specifies the range of buckets on which to calculate interim results.
//
func (f MlFlushJob) WithStart(v string) func(*MlFlushJobRequest) {
	return func(r *MlFlushJobRequest) {
		r.Start = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f MlFlushJob) WithPretty() func(*MlFlushJobRequest) {
	return func(r *MlFlushJobRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f MlFlushJob) WithHuman() func(*MlFlushJobRequest) {
	return func(r *MlFlushJobRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f MlFlushJob) WithErrorTrace() func(*MlFlushJobRequest) {
	return func(r *MlFlushJobRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f MlFlushJob) WithFilterPath(v ...string) func(*MlFlushJobRequest) {
	return func(r *MlFlushJobRequest) {
		r.FilterPath = v
	}
}
