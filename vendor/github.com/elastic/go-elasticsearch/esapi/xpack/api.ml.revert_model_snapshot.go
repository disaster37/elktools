// Code generated from specification version 7.0.0-SNAPSHOT: DO NOT EDIT

package xpack

import (
	"context"
	"io"
	"strconv"
	"strings"
)

func newMlRevertModelSnapshotFunc(t Transport) MlRevertModelSnapshot {
	return func(job_id string, snapshot_id string, o ...func(*MlRevertModelSnapshotRequest)) (*Response, error) {
		var r = MlRevertModelSnapshotRequest{JobID: job_id, SnapshotID: snapshot_id}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

//
// See full documentation at http://www.elastic.co/guide/en/elasticsearch/reference/current/ml-revert-snapshot.html.
//
type MlRevertModelSnapshot func(snapshot_id string, job_id string, o ...func(*MlRevertModelSnapshotRequest)) (*Response, error)

// MlRevertModelSnapshotRequest configures the Ml   Revert Model Snapshot API request.
//
type MlRevertModelSnapshotRequest struct {
	Body io.Reader

	JobID                    string
	SnapshotID               string
	DeleteInterveningResults *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r MlRevertModelSnapshotRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "POST"

	path.Grow(1 + len("_ml") + 1 + len("anomaly_detectors") + 1 + len(r.JobID) + 1 + len("model_snapshots") + 1 + len(r.SnapshotID) + 1 + len("_revert"))
	path.WriteString("/")
	path.WriteString("_ml")
	path.WriteString("/")
	path.WriteString("anomaly_detectors")
	path.WriteString("/")
	path.WriteString(r.JobID)
	path.WriteString("/")
	path.WriteString("model_snapshots")
	path.WriteString("/")
	path.WriteString(r.SnapshotID)
	path.WriteString("/")
	path.WriteString("_revert")

	params = make(map[string]string)

	if r.DeleteInterveningResults != nil {
		params["delete_intervening_results"] = strconv.FormatBool(*r.DeleteInterveningResults)
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
func (f MlRevertModelSnapshot) WithContext(v context.Context) func(*MlRevertModelSnapshotRequest) {
	return func(r *MlRevertModelSnapshotRequest) {
		r.ctx = v
	}
}

// WithBody - Reversion options.
//
func (f MlRevertModelSnapshot) WithBody(v io.Reader) func(*MlRevertModelSnapshotRequest) {
	return func(r *MlRevertModelSnapshotRequest) {
		r.Body = v
	}
}

// WithDeleteInterveningResults - should we reset the results back to the time of the snapshot?.
//
func (f MlRevertModelSnapshot) WithDeleteInterveningResults(v bool) func(*MlRevertModelSnapshotRequest) {
	return func(r *MlRevertModelSnapshotRequest) {
		r.DeleteInterveningResults = &v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f MlRevertModelSnapshot) WithPretty() func(*MlRevertModelSnapshotRequest) {
	return func(r *MlRevertModelSnapshotRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f MlRevertModelSnapshot) WithHuman() func(*MlRevertModelSnapshotRequest) {
	return func(r *MlRevertModelSnapshotRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f MlRevertModelSnapshot) WithErrorTrace() func(*MlRevertModelSnapshotRequest) {
	return func(r *MlRevertModelSnapshotRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f MlRevertModelSnapshot) WithFilterPath(v ...string) func(*MlRevertModelSnapshotRequest) {
	return func(r *MlRevertModelSnapshotRequest) {
		r.FilterPath = v
	}
}
