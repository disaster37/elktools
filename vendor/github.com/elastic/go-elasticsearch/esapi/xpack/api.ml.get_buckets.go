// Code generated from specification version 7.0.0-SNAPSHOT: DO NOT EDIT

package xpack

import (
	"context"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func newMlGetBucketsFunc(t Transport) MlGetBuckets {
	return func(job_id string, o ...func(*MlGetBucketsRequest)) (*Response, error) {
		var r = MlGetBucketsRequest{JobID: job_id}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

//
// See full documentation at http://www.elastic.co/guide/en/elasticsearch/reference/current/ml-get-bucket.html.
//
type MlGetBuckets func(job_id string, o ...func(*MlGetBucketsRequest)) (*Response, error)

// MlGetBucketsRequest configures the Ml  Get Buckets API request.
//
type MlGetBucketsRequest struct {
	Body io.Reader

	JobID          string
	Timestamp      string
	AnomalyScore   interface{}
	Desc           *bool
	End            string
	ExcludeInterim *bool
	Expand         *bool
	From           interface{}
	Size           interface{}
	Sort           string
	Start          string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r MlGetBucketsRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "GET"

	path.Grow(1 + len("_ml") + 1 + len("anomaly_detectors") + 1 + len(r.JobID) + 1 + len("results") + 1 + len("buckets") + 1 + len(r.Timestamp))
	path.WriteString("/")
	path.WriteString("_ml")
	path.WriteString("/")
	path.WriteString("anomaly_detectors")
	path.WriteString("/")
	path.WriteString(r.JobID)
	path.WriteString("/")
	path.WriteString("results")
	path.WriteString("/")
	path.WriteString("buckets")
	if r.Timestamp != "" {
		path.WriteString("/")
		path.WriteString(r.Timestamp)
	}

	params = make(map[string]string)

	if r.AnomalyScore != nil {
		params["anomaly_score"] = fmt.Sprintf("%v", r.AnomalyScore)
	}

	if r.Desc != nil {
		params["desc"] = strconv.FormatBool(*r.Desc)
	}

	if r.End != "" {
		params["end"] = r.End
	}

	if r.ExcludeInterim != nil {
		params["exclude_interim"] = strconv.FormatBool(*r.ExcludeInterim)
	}

	if r.Expand != nil {
		params["expand"] = strconv.FormatBool(*r.Expand)
	}

	if r.From != nil {
		params["from"] = fmt.Sprintf("%v", r.From)
	}

	if r.Size != nil {
		params["size"] = fmt.Sprintf("%v", r.Size)
	}

	if r.Sort != "" {
		params["sort"] = r.Sort
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
func (f MlGetBuckets) WithContext(v context.Context) func(*MlGetBucketsRequest) {
	return func(r *MlGetBucketsRequest) {
		r.ctx = v
	}
}

// WithTimestamp - the timestamp of the desired single bucket result.
//
func (f MlGetBuckets) WithTimestamp(v string) func(*MlGetBucketsRequest) {
	return func(r *MlGetBucketsRequest) {
		r.Timestamp = v
	}
}

// WithBody - Bucket selection details if not provided in URI.
//
func (f MlGetBuckets) WithBody(v io.Reader) func(*MlGetBucketsRequest) {
	return func(r *MlGetBucketsRequest) {
		r.Body = v
	}
}

// WithAnomalyScore - filter for the most anomalous buckets.
//
func (f MlGetBuckets) WithAnomalyScore(v interface{}) func(*MlGetBucketsRequest) {
	return func(r *MlGetBucketsRequest) {
		r.AnomalyScore = v
	}
}

// WithDesc - set the sort direction.
//
func (f MlGetBuckets) WithDesc(v bool) func(*MlGetBucketsRequest) {
	return func(r *MlGetBucketsRequest) {
		r.Desc = &v
	}
}

// WithEnd - end time filter for buckets.
//
func (f MlGetBuckets) WithEnd(v string) func(*MlGetBucketsRequest) {
	return func(r *MlGetBucketsRequest) {
		r.End = v
	}
}

// WithExcludeInterim - exclude interim results.
//
func (f MlGetBuckets) WithExcludeInterim(v bool) func(*MlGetBucketsRequest) {
	return func(r *MlGetBucketsRequest) {
		r.ExcludeInterim = &v
	}
}

// WithExpand - include anomaly records.
//
func (f MlGetBuckets) WithExpand(v bool) func(*MlGetBucketsRequest) {
	return func(r *MlGetBucketsRequest) {
		r.Expand = &v
	}
}

// WithFrom - skips a number of buckets.
//
func (f MlGetBuckets) WithFrom(v interface{}) func(*MlGetBucketsRequest) {
	return func(r *MlGetBucketsRequest) {
		r.From = v
	}
}

// WithSize - specifies a max number of buckets to get.
//
func (f MlGetBuckets) WithSize(v interface{}) func(*MlGetBucketsRequest) {
	return func(r *MlGetBucketsRequest) {
		r.Size = v
	}
}

// WithSort - sort buckets by a particular field.
//
func (f MlGetBuckets) WithSort(v string) func(*MlGetBucketsRequest) {
	return func(r *MlGetBucketsRequest) {
		r.Sort = v
	}
}

// WithStart - start time filter for buckets.
//
func (f MlGetBuckets) WithStart(v string) func(*MlGetBucketsRequest) {
	return func(r *MlGetBucketsRequest) {
		r.Start = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f MlGetBuckets) WithPretty() func(*MlGetBucketsRequest) {
	return func(r *MlGetBucketsRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f MlGetBuckets) WithHuman() func(*MlGetBucketsRequest) {
	return func(r *MlGetBucketsRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f MlGetBuckets) WithErrorTrace() func(*MlGetBucketsRequest) {
	return func(r *MlGetBucketsRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f MlGetBuckets) WithFilterPath(v ...string) func(*MlGetBucketsRequest) {
	return func(r *MlGetBucketsRequest) {
		r.FilterPath = v
	}
}
