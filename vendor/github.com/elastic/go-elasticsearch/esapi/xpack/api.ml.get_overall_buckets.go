// Code generated from specification version 7.0.0-SNAPSHOT: DO NOT EDIT

package xpack

import (
	"context"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func newMlGetOverallBucketsFunc(t Transport) MlGetOverallBuckets {
	return func(job_id string, o ...func(*MlGetOverallBucketsRequest)) (*Response, error) {
		var r = MlGetOverallBucketsRequest{JobID: job_id}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

//
// See full documentation at http://www.elastic.co/guide/en/elasticsearch/reference/current/ml-get-overall-buckets.html.
//
type MlGetOverallBuckets func(job_id string, o ...func(*MlGetOverallBucketsRequest)) (*Response, error)

// MlGetOverallBucketsRequest configures the Ml   Get Overall Buckets API request.
//
type MlGetOverallBucketsRequest struct {
	Body io.Reader

	JobID          string
	AllowNoJobs    *bool
	BucketSpan     string
	End            string
	ExcludeInterim *bool
	OverallScore   interface{}
	Start          string
	TopN           interface{}

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r MlGetOverallBucketsRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "GET"

	path.Grow(1 + len("_ml") + 1 + len("anomaly_detectors") + 1 + len(r.JobID) + 1 + len("results") + 1 + len("overall_buckets"))
	path.WriteString("/")
	path.WriteString("_ml")
	path.WriteString("/")
	path.WriteString("anomaly_detectors")
	path.WriteString("/")
	path.WriteString(r.JobID)
	path.WriteString("/")
	path.WriteString("results")
	path.WriteString("/")
	path.WriteString("overall_buckets")

	params = make(map[string]string)

	if r.AllowNoJobs != nil {
		params["allow_no_jobs"] = strconv.FormatBool(*r.AllowNoJobs)
	}

	if r.BucketSpan != "" {
		params["bucket_span"] = r.BucketSpan
	}

	if r.End != "" {
		params["end"] = r.End
	}

	if r.ExcludeInterim != nil {
		params["exclude_interim"] = strconv.FormatBool(*r.ExcludeInterim)
	}

	if r.OverallScore != nil {
		params["overall_score"] = fmt.Sprintf("%v", r.OverallScore)
	}

	if r.Start != "" {
		params["start"] = r.Start
	}

	if r.TopN != nil {
		params["top_n"] = fmt.Sprintf("%v", r.TopN)
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
func (f MlGetOverallBuckets) WithContext(v context.Context) func(*MlGetOverallBucketsRequest) {
	return func(r *MlGetOverallBucketsRequest) {
		r.ctx = v
	}
}

// WithBody - Overall bucket selection details if not provided in URI.
//
func (f MlGetOverallBuckets) WithBody(v io.Reader) func(*MlGetOverallBucketsRequest) {
	return func(r *MlGetOverallBucketsRequest) {
		r.Body = v
	}
}

// WithAllowNoJobs - whether to ignore if a wildcard expression matches no jobs. (this includes `_all` string or when no jobs have been specified).
//
func (f MlGetOverallBuckets) WithAllowNoJobs(v bool) func(*MlGetOverallBucketsRequest) {
	return func(r *MlGetOverallBucketsRequest) {
		r.AllowNoJobs = &v
	}
}

// WithBucketSpan - the span of the overall buckets. defaults to the longest job bucket_span.
//
func (f MlGetOverallBuckets) WithBucketSpan(v string) func(*MlGetOverallBucketsRequest) {
	return func(r *MlGetOverallBucketsRequest) {
		r.BucketSpan = v
	}
}

// WithEnd - returns overall buckets with timestamps earlier than this time.
//
func (f MlGetOverallBuckets) WithEnd(v string) func(*MlGetOverallBucketsRequest) {
	return func(r *MlGetOverallBucketsRequest) {
		r.End = v
	}
}

// WithExcludeInterim - if true overall buckets that include interim buckets will be excluded.
//
func (f MlGetOverallBuckets) WithExcludeInterim(v bool) func(*MlGetOverallBucketsRequest) {
	return func(r *MlGetOverallBucketsRequest) {
		r.ExcludeInterim = &v
	}
}

// WithOverallScore - returns overall buckets with overall scores higher than this value.
//
func (f MlGetOverallBuckets) WithOverallScore(v interface{}) func(*MlGetOverallBucketsRequest) {
	return func(r *MlGetOverallBucketsRequest) {
		r.OverallScore = v
	}
}

// WithStart - returns overall buckets with timestamps after this time.
//
func (f MlGetOverallBuckets) WithStart(v string) func(*MlGetOverallBucketsRequest) {
	return func(r *MlGetOverallBucketsRequest) {
		r.Start = v
	}
}

// WithTopN - the number of top job bucket scores to be used in the overall_score calculation.
//
func (f MlGetOverallBuckets) WithTopN(v interface{}) func(*MlGetOverallBucketsRequest) {
	return func(r *MlGetOverallBucketsRequest) {
		r.TopN = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f MlGetOverallBuckets) WithPretty() func(*MlGetOverallBucketsRequest) {
	return func(r *MlGetOverallBucketsRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f MlGetOverallBuckets) WithHuman() func(*MlGetOverallBucketsRequest) {
	return func(r *MlGetOverallBucketsRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f MlGetOverallBuckets) WithErrorTrace() func(*MlGetOverallBucketsRequest) {
	return func(r *MlGetOverallBucketsRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f MlGetOverallBuckets) WithFilterPath(v ...string) func(*MlGetOverallBucketsRequest) {
	return func(r *MlGetOverallBucketsRequest) {
		r.FilterPath = v
	}
}
