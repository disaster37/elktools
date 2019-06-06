// Code generated from specification version 7.0.0-SNAPSHOT: DO NOT EDIT

package xpack

import (
	"context"
	"io"
	"strings"
)

func newMlPutDatafeedFunc(t Transport) MlPutDatafeed {
	return func(body io.Reader, datafeed_id string, o ...func(*MlPutDatafeedRequest)) (*Response, error) {
		var r = MlPutDatafeedRequest{Body: body, DatafeedID: datafeed_id}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

//
// See full documentation at http://www.elastic.co/guide/en/elasticsearch/reference/current/ml-put-datafeed.html.
//
type MlPutDatafeed func(body io.Reader, datafeed_id string, o ...func(*MlPutDatafeedRequest)) (*Response, error)

// MlPutDatafeedRequest configures the Ml  Put Datafeed API request.
//
type MlPutDatafeedRequest struct {
	Body io.Reader

	DatafeedID string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r MlPutDatafeedRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "PUT"

	path.Grow(1 + len("_ml") + 1 + len("datafeeds") + 1 + len(r.DatafeedID))
	path.WriteString("/")
	path.WriteString("_ml")
	path.WriteString("/")
	path.WriteString("datafeeds")
	path.WriteString("/")
	path.WriteString(r.DatafeedID)

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
func (f MlPutDatafeed) WithContext(v context.Context) func(*MlPutDatafeedRequest) {
	return func(r *MlPutDatafeedRequest) {
		r.ctx = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f MlPutDatafeed) WithPretty() func(*MlPutDatafeedRequest) {
	return func(r *MlPutDatafeedRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f MlPutDatafeed) WithHuman() func(*MlPutDatafeedRequest) {
	return func(r *MlPutDatafeedRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f MlPutDatafeed) WithErrorTrace() func(*MlPutDatafeedRequest) {
	return func(r *MlPutDatafeedRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f MlPutDatafeed) WithFilterPath(v ...string) func(*MlPutDatafeedRequest) {
	return func(r *MlPutDatafeedRequest) {
		r.FilterPath = v
	}
}
