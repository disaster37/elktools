// Code generated from specification version 7.0.0-SNAPSHOT: DO NOT EDIT

package xpack

import (
	"context"
	"strconv"
	"strings"
)

func newMlDeleteDatafeedFunc(t Transport) MlDeleteDatafeed {
	return func(datafeed_id string, o ...func(*MlDeleteDatafeedRequest)) (*Response, error) {
		var r = MlDeleteDatafeedRequest{DatafeedID: datafeed_id}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

//
// See full documentation at http://www.elastic.co/guide/en/elasticsearch/reference/current/ml-delete-datafeed.html.
//
type MlDeleteDatafeed func(datafeed_id string, o ...func(*MlDeleteDatafeedRequest)) (*Response, error)

// MlDeleteDatafeedRequest configures the Ml  Delete Datafeed API request.
//
type MlDeleteDatafeedRequest struct {
	DatafeedID string
	Force      *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r MlDeleteDatafeedRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "DELETE"

	path.Grow(1 + len("_ml") + 1 + len("datafeeds") + 1 + len(r.DatafeedID))
	path.WriteString("/")
	path.WriteString("_ml")
	path.WriteString("/")
	path.WriteString("datafeeds")
	path.WriteString("/")
	path.WriteString(r.DatafeedID)

	params = make(map[string]string)

	if r.Force != nil {
		params["force"] = strconv.FormatBool(*r.Force)
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
func (f MlDeleteDatafeed) WithContext(v context.Context) func(*MlDeleteDatafeedRequest) {
	return func(r *MlDeleteDatafeedRequest) {
		r.ctx = v
	}
}

// WithForce - true if the datafeed should be forcefully deleted.
//
func (f MlDeleteDatafeed) WithForce(v bool) func(*MlDeleteDatafeedRequest) {
	return func(r *MlDeleteDatafeedRequest) {
		r.Force = &v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f MlDeleteDatafeed) WithPretty() func(*MlDeleteDatafeedRequest) {
	return func(r *MlDeleteDatafeedRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f MlDeleteDatafeed) WithHuman() func(*MlDeleteDatafeedRequest) {
	return func(r *MlDeleteDatafeedRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f MlDeleteDatafeed) WithErrorTrace() func(*MlDeleteDatafeedRequest) {
	return func(r *MlDeleteDatafeedRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f MlDeleteDatafeed) WithFilterPath(v ...string) func(*MlDeleteDatafeedRequest) {
	return func(r *MlDeleteDatafeedRequest) {
		r.FilterPath = v
	}
}
