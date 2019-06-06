// Code generated from specification version 7.0.0-SNAPSHOT: DO NOT EDIT

package xpack

import (
	"context"
	"strings"
)

func newMlPreviewDatafeedFunc(t Transport) MlPreviewDatafeed {
	return func(datafeed_id string, o ...func(*MlPreviewDatafeedRequest)) (*Response, error) {
		var r = MlPreviewDatafeedRequest{DatafeedID: datafeed_id}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

//
// See full documentation at http://www.elastic.co/guide/en/elasticsearch/reference/current/ml-preview-datafeed.html.
//
type MlPreviewDatafeed func(datafeed_id string, o ...func(*MlPreviewDatafeedRequest)) (*Response, error)

// MlPreviewDatafeedRequest configures the Ml  Preview Datafeed API request.
//
type MlPreviewDatafeedRequest struct {
	DatafeedID string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r MlPreviewDatafeedRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "GET"

	path.Grow(1 + len("_ml") + 1 + len("datafeeds") + 1 + len(r.DatafeedID) + 1 + len("_preview"))
	path.WriteString("/")
	path.WriteString("_ml")
	path.WriteString("/")
	path.WriteString("datafeeds")
	path.WriteString("/")
	path.WriteString(r.DatafeedID)
	path.WriteString("/")
	path.WriteString("_preview")

	params = make(map[string]string)

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
func (f MlPreviewDatafeed) WithContext(v context.Context) func(*MlPreviewDatafeedRequest) {
	return func(r *MlPreviewDatafeedRequest) {
		r.ctx = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f MlPreviewDatafeed) WithPretty() func(*MlPreviewDatafeedRequest) {
	return func(r *MlPreviewDatafeedRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f MlPreviewDatafeed) WithHuman() func(*MlPreviewDatafeedRequest) {
	return func(r *MlPreviewDatafeedRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f MlPreviewDatafeed) WithErrorTrace() func(*MlPreviewDatafeedRequest) {
	return func(r *MlPreviewDatafeedRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f MlPreviewDatafeed) WithFilterPath(v ...string) func(*MlPreviewDatafeedRequest) {
	return func(r *MlPreviewDatafeedRequest) {
		r.FilterPath = v
	}
}
