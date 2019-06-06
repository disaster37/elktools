// Code generated from specification version 7.0.0-SNAPSHOT: DO NOT EDIT

package xpack

import (
	"context"
	"strconv"
	"strings"
	"time"
)

func newMlStopDatafeedFunc(t Transport) MlStopDatafeed {
	return func(datafeed_id string, o ...func(*MlStopDatafeedRequest)) (*Response, error) {
		var r = MlStopDatafeedRequest{DatafeedID: datafeed_id}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

//
// See full documentation at http://www.elastic.co/guide/en/elasticsearch/reference/current/ml-stop-datafeed.html.
//
type MlStopDatafeed func(datafeed_id string, o ...func(*MlStopDatafeedRequest)) (*Response, error)

// MlStopDatafeedRequest configures the Ml  Stop Datafeed API request.
//
type MlStopDatafeedRequest struct {
	DatafeedID       string
	AllowNoDatafeeds *bool
	Force            *bool
	Timeout          time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r MlStopDatafeedRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "POST"

	path.Grow(1 + len("_ml") + 1 + len("datafeeds") + 1 + len(r.DatafeedID) + 1 + len("_stop"))
	path.WriteString("/")
	path.WriteString("_ml")
	path.WriteString("/")
	path.WriteString("datafeeds")
	path.WriteString("/")
	path.WriteString(r.DatafeedID)
	path.WriteString("/")
	path.WriteString("_stop")

	params = make(map[string]string)

	if r.AllowNoDatafeeds != nil {
		params["allow_no_datafeeds"] = strconv.FormatBool(*r.AllowNoDatafeeds)
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
func (f MlStopDatafeed) WithContext(v context.Context) func(*MlStopDatafeedRequest) {
	return func(r *MlStopDatafeedRequest) {
		r.ctx = v
	}
}

// WithAllowNoDatafeeds - whether to ignore if a wildcard expression matches no datafeeds. (this includes `_all` string or when no datafeeds have been specified).
//
func (f MlStopDatafeed) WithAllowNoDatafeeds(v bool) func(*MlStopDatafeedRequest) {
	return func(r *MlStopDatafeedRequest) {
		r.AllowNoDatafeeds = &v
	}
}

// WithForce - true if the datafeed should be forcefully stopped..
//
func (f MlStopDatafeed) WithForce(v bool) func(*MlStopDatafeedRequest) {
	return func(r *MlStopDatafeedRequest) {
		r.Force = &v
	}
}

// WithTimeout - controls the time to wait until a datafeed has stopped. default to 20 seconds.
//
func (f MlStopDatafeed) WithTimeout(v time.Duration) func(*MlStopDatafeedRequest) {
	return func(r *MlStopDatafeedRequest) {
		r.Timeout = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f MlStopDatafeed) WithPretty() func(*MlStopDatafeedRequest) {
	return func(r *MlStopDatafeedRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f MlStopDatafeed) WithHuman() func(*MlStopDatafeedRequest) {
	return func(r *MlStopDatafeedRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f MlStopDatafeed) WithErrorTrace() func(*MlStopDatafeedRequest) {
	return func(r *MlStopDatafeedRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f MlStopDatafeed) WithFilterPath(v ...string) func(*MlStopDatafeedRequest) {
	return func(r *MlStopDatafeedRequest) {
		r.FilterPath = v
	}
}
