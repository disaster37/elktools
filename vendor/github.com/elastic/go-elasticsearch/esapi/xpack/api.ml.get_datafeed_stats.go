// Code generated from specification version 7.0.0-SNAPSHOT: DO NOT EDIT

package xpack

import (
	"context"
	"strconv"
	"strings"
)

func newMlGetDatafeedStatsFunc(t Transport) MlGetDatafeedStats {
	return func(o ...func(*MlGetDatafeedStatsRequest)) (*Response, error) {
		var r = MlGetDatafeedStatsRequest{}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

//
// See full documentation at http://www.elastic.co/guide/en/elasticsearch/reference/current/ml-get-datafeed-stats.html.
//
type MlGetDatafeedStats func(o ...func(*MlGetDatafeedStatsRequest)) (*Response, error)

// MlGetDatafeedStatsRequest configures the Ml   Get Datafeed Stats API request.
//
type MlGetDatafeedStatsRequest struct {
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
func (r MlGetDatafeedStatsRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "GET"

	path.Grow(1 + len("_ml") + 1 + len("datafeeds") + 1 + len(r.DatafeedID) + 1 + len("_stats"))
	path.WriteString("/")
	path.WriteString("_ml")
	path.WriteString("/")
	path.WriteString("datafeeds")
	if r.DatafeedID != "" {
		path.WriteString("/")
		path.WriteString(r.DatafeedID)
	}
	path.WriteString("/")
	path.WriteString("_stats")

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
func (f MlGetDatafeedStats) WithContext(v context.Context) func(*MlGetDatafeedStatsRequest) {
	return func(r *MlGetDatafeedStatsRequest) {
		r.ctx = v
	}
}

// WithDatafeedID - the ID of the datafeeds stats to fetch.
//
func (f MlGetDatafeedStats) WithDatafeedID(v string) func(*MlGetDatafeedStatsRequest) {
	return func(r *MlGetDatafeedStatsRequest) {
		r.DatafeedID = v
	}
}

// WithAllowNoDatafeeds - whether to ignore if a wildcard expression matches no datafeeds. (this includes `_all` string or when no datafeeds have been specified).
//
func (f MlGetDatafeedStats) WithAllowNoDatafeeds(v bool) func(*MlGetDatafeedStatsRequest) {
	return func(r *MlGetDatafeedStatsRequest) {
		r.AllowNoDatafeeds = &v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f MlGetDatafeedStats) WithPretty() func(*MlGetDatafeedStatsRequest) {
	return func(r *MlGetDatafeedStatsRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f MlGetDatafeedStats) WithHuman() func(*MlGetDatafeedStatsRequest) {
	return func(r *MlGetDatafeedStatsRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f MlGetDatafeedStats) WithErrorTrace() func(*MlGetDatafeedStatsRequest) {
	return func(r *MlGetDatafeedStatsRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f MlGetDatafeedStats) WithFilterPath(v ...string) func(*MlGetDatafeedStatsRequest) {
	return func(r *MlGetDatafeedStatsRequest) {
		r.FilterPath = v
	}
}
