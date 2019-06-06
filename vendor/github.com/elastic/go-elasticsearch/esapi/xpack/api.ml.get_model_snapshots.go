// Code generated from specification version 7.0.0-SNAPSHOT: DO NOT EDIT

package xpack

import (
	"context"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func newMlGetModelSnapshotsFunc(t Transport) MlGetModelSnapshots {
	return func(job_id string, o ...func(*MlGetModelSnapshotsRequest)) (*Response, error) {
		var r = MlGetModelSnapshotsRequest{JobID: job_id}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

//
// See full documentation at http://www.elastic.co/guide/en/elasticsearch/reference/current/ml-get-snapshot.html.
//
type MlGetModelSnapshots func(job_id string, o ...func(*MlGetModelSnapshotsRequest)) (*Response, error)

// MlGetModelSnapshotsRequest configures the Ml   Get Model Snapshots API request.
//
type MlGetModelSnapshotsRequest struct {
	Body io.Reader

	JobID      string
	SnapshotID string
	Desc       *bool
	End        interface{}
	From       interface{}
	Size       interface{}
	Sort       string
	Start      interface{}

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r MlGetModelSnapshotsRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "GET"

	path.Grow(1 + len("_ml") + 1 + len("anomaly_detectors") + 1 + len(r.JobID) + 1 + len("model_snapshots") + 1 + len(r.SnapshotID))
	path.WriteString("/")
	path.WriteString("_ml")
	path.WriteString("/")
	path.WriteString("anomaly_detectors")
	path.WriteString("/")
	path.WriteString(r.JobID)
	path.WriteString("/")
	path.WriteString("model_snapshots")
	if r.SnapshotID != "" {
		path.WriteString("/")
		path.WriteString(r.SnapshotID)
	}

	params = make(map[string]string)

	if r.Desc != nil {
		params["desc"] = strconv.FormatBool(*r.Desc)
	}

	if r.End != nil {
		params["end"] = fmt.Sprintf("%v", r.End)
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

	if r.Start != nil {
		params["start"] = fmt.Sprintf("%v", r.Start)
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
func (f MlGetModelSnapshots) WithContext(v context.Context) func(*MlGetModelSnapshotsRequest) {
	return func(r *MlGetModelSnapshotsRequest) {
		r.ctx = v
	}
}

// WithSnapshotID - the ID of the snapshot to fetch.
//
func (f MlGetModelSnapshots) WithSnapshotID(v string) func(*MlGetModelSnapshotsRequest) {
	return func(r *MlGetModelSnapshotsRequest) {
		r.SnapshotID = v
	}
}

// WithBody - Model snapshot selection criteria.
//
func (f MlGetModelSnapshots) WithBody(v io.Reader) func(*MlGetModelSnapshotsRequest) {
	return func(r *MlGetModelSnapshotsRequest) {
		r.Body = v
	}
}

// WithDesc - true if the results should be sorted in descending order.
//
func (f MlGetModelSnapshots) WithDesc(v bool) func(*MlGetModelSnapshotsRequest) {
	return func(r *MlGetModelSnapshotsRequest) {
		r.Desc = &v
	}
}

// WithEnd - the filter 'end' query parameter.
//
func (f MlGetModelSnapshots) WithEnd(v interface{}) func(*MlGetModelSnapshotsRequest) {
	return func(r *MlGetModelSnapshotsRequest) {
		r.End = v
	}
}

// WithFrom - skips a number of documents.
//
func (f MlGetModelSnapshots) WithFrom(v interface{}) func(*MlGetModelSnapshotsRequest) {
	return func(r *MlGetModelSnapshotsRequest) {
		r.From = v
	}
}

// WithSize - the default number of documents returned in queries as a string..
//
func (f MlGetModelSnapshots) WithSize(v interface{}) func(*MlGetModelSnapshotsRequest) {
	return func(r *MlGetModelSnapshotsRequest) {
		r.Size = v
	}
}

// WithSort - name of the field to sort on.
//
func (f MlGetModelSnapshots) WithSort(v string) func(*MlGetModelSnapshotsRequest) {
	return func(r *MlGetModelSnapshotsRequest) {
		r.Sort = v
	}
}

// WithStart - the filter 'start' query parameter.
//
func (f MlGetModelSnapshots) WithStart(v interface{}) func(*MlGetModelSnapshotsRequest) {
	return func(r *MlGetModelSnapshotsRequest) {
		r.Start = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f MlGetModelSnapshots) WithPretty() func(*MlGetModelSnapshotsRequest) {
	return func(r *MlGetModelSnapshotsRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f MlGetModelSnapshots) WithHuman() func(*MlGetModelSnapshotsRequest) {
	return func(r *MlGetModelSnapshotsRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f MlGetModelSnapshots) WithErrorTrace() func(*MlGetModelSnapshotsRequest) {
	return func(r *MlGetModelSnapshotsRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f MlGetModelSnapshots) WithFilterPath(v ...string) func(*MlGetModelSnapshotsRequest) {
	return func(r *MlGetModelSnapshotsRequest) {
		r.FilterPath = v
	}
}
