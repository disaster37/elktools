// Code generated from specification version 7.0.0-SNAPSHOT: DO NOT EDIT

package xpack

import (
	"context"
	"strconv"
	"strings"
)

func newMlGetDatafeedsFunc(t Transport) MlGetDatafeeds {
	return func(o ...func(*MlGetDatafeedsRequest)) (*Response, error) {
		var r = MlGetDatafeedsRequest{}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

//
// See full documentation at http://www.elastic.co/guide/en/elasticsearch/reference/current/ml-get-datafeed.html.
//
type MlGetDatafeeds func(o ...func(*MlGetDatafeedsRequest)) (*Response, error)

// MlGetDatafeedsRequest configures the Ml  Get Datafeeds API request.
//
type MlGetDatafeedsRequest struct {
	DatafeedID       string
	AllowNoDatafeeds *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r MlGetDatafeedsRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "GET"

	path.Grow(1 + len("_ml") + 1 + len("datafeeds") + 1 + len(r.DatafeedID))
	path.WriteString("/")
	path.WriteString("_ml")
	path.WriteString("/")
	path.WriteString("datafeeds")
	if r.DatafeedID != "" {
		path.WriteString("/")
		path.WriteString(r.DatafeedID)
	}

	params = make(map[string]string)

	if r.AllowNoDatafeeds != nil {
		params["allow_no_datafeeds"] = strconv.FormatBool(*r.AllowNoDatafeeds)
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
func (f MlGetDatafeeds) WithContext(v context.Context) func(*MlGetDatafeedsRequest) {
	return func(r *MlGetDatafeedsRequest) {
		r.ctx = v
	}
}

// WithDatafeedID - the ID of the datafeeds to fetch.
//
func (f MlGetDatafeeds) WithDatafeedID(v string) func(*MlGetDatafeedsRequest) {
	return func(r *MlGetDatafeedsRequest) {
		r.DatafeedID = v
	}
}

// WithAllowNoDatafeeds - whether to ignore if a wildcard expression matches no datafeeds. (this includes `_all` string or when no datafeeds have been specified).
//
func (f MlGetDatafeeds) WithAllowNoDatafeeds(v bool) func(*MlGetDatafeedsRequest) {
	return func(r *MlGetDatafeedsRequest) {
		r.AllowNoDatafeeds = &v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f MlGetDatafeeds) WithPretty() func(*MlGetDatafeedsRequest) {
	return func(r *MlGetDatafeedsRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f MlGetDatafeeds) WithHuman() func(*MlGetDatafeedsRequest) {
	return func(r *MlGetDatafeedsRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f MlGetDatafeeds) WithErrorTrace() func(*MlGetDatafeedsRequest) {
	return func(r *MlGetDatafeedsRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f MlGetDatafeeds) WithFilterPath(v ...string) func(*MlGetDatafeedsRequest) {
	return func(r *MlGetDatafeedsRequest) {
		r.FilterPath = v
	}
}
