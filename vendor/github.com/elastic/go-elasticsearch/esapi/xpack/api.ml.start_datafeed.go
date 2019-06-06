// Code generated from specification version 7.0.0-SNAPSHOT: DO NOT EDIT

package xpack

import (
	"context"
	"io"
	"strings"
	"time"
)

func newMlStartDatafeedFunc(t Transport) MlStartDatafeed {
	return func(datafeed_id string, o ...func(*MlStartDatafeedRequest)) (*Response, error) {
		var r = MlStartDatafeedRequest{DatafeedID: datafeed_id}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

//
// See full documentation at http://www.elastic.co/guide/en/elasticsearch/reference/current/ml-start-datafeed.html.
//
type MlStartDatafeed func(datafeed_id string, o ...func(*MlStartDatafeedRequest)) (*Response, error)

// MlStartDatafeedRequest configures the Ml  Start Datafeed API request.
//
type MlStartDatafeedRequest struct {
	Body io.Reader

	DatafeedID string
	End        string
	Start      string
	Timeout    time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r MlStartDatafeedRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "POST"

	path.Grow(1 + len("_ml") + 1 + len("datafeeds") + 1 + len(r.DatafeedID) + 1 + len("_start"))
	path.WriteString("/")
	path.WriteString("_ml")
	path.WriteString("/")
	path.WriteString("datafeeds")
	path.WriteString("/")
	path.WriteString(r.DatafeedID)
	path.WriteString("/")
	path.WriteString("_start")

	params = make(map[string]string)

	if r.End != "" {
		params["end"] = r.End
	}

	if r.Start != "" {
		params["start"] = r.Start
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
func (f MlStartDatafeed) WithContext(v context.Context) func(*MlStartDatafeedRequest) {
	return func(r *MlStartDatafeedRequest) {
		r.ctx = v
	}
}

// WithBody - The start datafeed parameters.
//
func (f MlStartDatafeed) WithBody(v io.Reader) func(*MlStartDatafeedRequest) {
	return func(r *MlStartDatafeedRequest) {
		r.Body = v
	}
}

// WithEnd - the end time when the datafeed should stop. when not set, the datafeed continues in real time.
//
func (f MlStartDatafeed) WithEnd(v string) func(*MlStartDatafeedRequest) {
	return func(r *MlStartDatafeedRequest) {
		r.End = v
	}
}

// WithStart - the start time from where the datafeed should begin.
//
func (f MlStartDatafeed) WithStart(v string) func(*MlStartDatafeedRequest) {
	return func(r *MlStartDatafeedRequest) {
		r.Start = v
	}
}

// WithTimeout - controls the time to wait until a datafeed has started. default to 20 seconds.
//
func (f MlStartDatafeed) WithTimeout(v time.Duration) func(*MlStartDatafeedRequest) {
	return func(r *MlStartDatafeedRequest) {
		r.Timeout = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f MlStartDatafeed) WithPretty() func(*MlStartDatafeedRequest) {
	return func(r *MlStartDatafeedRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f MlStartDatafeed) WithHuman() func(*MlStartDatafeedRequest) {
	return func(r *MlStartDatafeedRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f MlStartDatafeed) WithErrorTrace() func(*MlStartDatafeedRequest) {
	return func(r *MlStartDatafeedRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f MlStartDatafeed) WithFilterPath(v ...string) func(*MlStartDatafeedRequest) {
	return func(r *MlStartDatafeedRequest) {
		r.FilterPath = v
	}
}
