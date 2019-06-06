// Code generated from specification version 7.0.0-SNAPSHOT: DO NOT EDIT

package xpack

import (
	"context"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func newMlGetRecordsFunc(t Transport) MlGetRecords {
	return func(job_id string, o ...func(*MlGetRecordsRequest)) (*Response, error) {
		var r = MlGetRecordsRequest{JobID: job_id}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

//
// See full documentation at http://www.elastic.co/guide/en/elasticsearch/reference/current/ml-get-record.html.
//
type MlGetRecords func(job_id string, o ...func(*MlGetRecordsRequest)) (*Response, error)

// MlGetRecordsRequest configures the Ml  Get Records API request.
//
type MlGetRecordsRequest struct {
	Body io.Reader

	JobID          string
	Desc           *bool
	End            string
	ExcludeInterim *bool
	From           interface{}
	RecordScore    interface{}
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
func (r MlGetRecordsRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "GET"

	path.Grow(1 + len("_ml") + 1 + len("anomaly_detectors") + 1 + len(r.JobID) + 1 + len("results") + 1 + len("records"))
	path.WriteString("/")
	path.WriteString("_ml")
	path.WriteString("/")
	path.WriteString("anomaly_detectors")
	path.WriteString("/")
	path.WriteString(r.JobID)
	path.WriteString("/")
	path.WriteString("results")
	path.WriteString("/")
	path.WriteString("records")

	params = make(map[string]string)

	if r.Desc != nil {
		params["desc"] = strconv.FormatBool(*r.Desc)
	}

	if r.End != "" {
		params["end"] = r.End
	}

	if r.ExcludeInterim != nil {
		params["exclude_interim"] = strconv.FormatBool(*r.ExcludeInterim)
	}

	if r.From != nil {
		params["from"] = fmt.Sprintf("%v", r.From)
	}

	if r.RecordScore != nil {
		params["record_score"] = fmt.Sprintf("%v", r.RecordScore)
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
func (f MlGetRecords) WithContext(v context.Context) func(*MlGetRecordsRequest) {
	return func(r *MlGetRecordsRequest) {
		r.ctx = v
	}
}

// WithBody - Record selection criteria.
//
func (f MlGetRecords) WithBody(v io.Reader) func(*MlGetRecordsRequest) {
	return func(r *MlGetRecordsRequest) {
		r.Body = v
	}
}

// WithDesc - set the sort direction.
//
func (f MlGetRecords) WithDesc(v bool) func(*MlGetRecordsRequest) {
	return func(r *MlGetRecordsRequest) {
		r.Desc = &v
	}
}

// WithEnd - end time filter for records.
//
func (f MlGetRecords) WithEnd(v string) func(*MlGetRecordsRequest) {
	return func(r *MlGetRecordsRequest) {
		r.End = v
	}
}

// WithExcludeInterim - exclude interim results.
//
func (f MlGetRecords) WithExcludeInterim(v bool) func(*MlGetRecordsRequest) {
	return func(r *MlGetRecordsRequest) {
		r.ExcludeInterim = &v
	}
}

// WithFrom - skips a number of records.
//
func (f MlGetRecords) WithFrom(v interface{}) func(*MlGetRecordsRequest) {
	return func(r *MlGetRecordsRequest) {
		r.From = v
	}
}

// WithRecordScore - .
//
func (f MlGetRecords) WithRecordScore(v interface{}) func(*MlGetRecordsRequest) {
	return func(r *MlGetRecordsRequest) {
		r.RecordScore = v
	}
}

// WithSize - specifies a max number of records to get.
//
func (f MlGetRecords) WithSize(v interface{}) func(*MlGetRecordsRequest) {
	return func(r *MlGetRecordsRequest) {
		r.Size = v
	}
}

// WithSort - sort records by a particular field.
//
func (f MlGetRecords) WithSort(v string) func(*MlGetRecordsRequest) {
	return func(r *MlGetRecordsRequest) {
		r.Sort = v
	}
}

// WithStart - start time filter for records.
//
func (f MlGetRecords) WithStart(v string) func(*MlGetRecordsRequest) {
	return func(r *MlGetRecordsRequest) {
		r.Start = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f MlGetRecords) WithPretty() func(*MlGetRecordsRequest) {
	return func(r *MlGetRecordsRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f MlGetRecords) WithHuman() func(*MlGetRecordsRequest) {
	return func(r *MlGetRecordsRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f MlGetRecords) WithErrorTrace() func(*MlGetRecordsRequest) {
	return func(r *MlGetRecordsRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f MlGetRecords) WithFilterPath(v ...string) func(*MlGetRecordsRequest) {
	return func(r *MlGetRecordsRequest) {
		r.FilterPath = v
	}
}
