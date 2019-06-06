// Code generated from specification version 7.0.0-SNAPSHOT: DO NOT EDIT

package xpack

import (
	"context"
	"fmt"
	"strings"
)

func newMlGetFiltersFunc(t Transport) MlGetFilters {
	return func(o ...func(*MlGetFiltersRequest)) (*Response, error) {
		var r = MlGetFiltersRequest{}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

//
type MlGetFilters func(o ...func(*MlGetFiltersRequest)) (*Response, error)

// MlGetFiltersRequest configures the Ml  Get Filters API request.
//
type MlGetFiltersRequest struct {
	FilterID string
	From     interface{}
	Size     interface{}

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r MlGetFiltersRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "GET"

	path.Grow(1 + len("_ml") + 1 + len("filters") + 1 + len(r.FilterID))
	path.WriteString("/")
	path.WriteString("_ml")
	path.WriteString("/")
	path.WriteString("filters")
	if r.FilterID != "" {
		path.WriteString("/")
		path.WriteString(r.FilterID)
	}

	params = make(map[string]string)

	if r.From != nil {
		params["from"] = fmt.Sprintf("%v", r.From)
	}

	if r.Size != nil {
		params["size"] = fmt.Sprintf("%v", r.Size)
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
func (f MlGetFilters) WithContext(v context.Context) func(*MlGetFiltersRequest) {
	return func(r *MlGetFiltersRequest) {
		r.ctx = v
	}
}

// WithFilterID - the ID of the filter to fetch.
//
func (f MlGetFilters) WithFilterID(v string) func(*MlGetFiltersRequest) {
	return func(r *MlGetFiltersRequest) {
		r.FilterID = v
	}
}

// WithFrom - skips a number of filters.
//
func (f MlGetFilters) WithFrom(v interface{}) func(*MlGetFiltersRequest) {
	return func(r *MlGetFiltersRequest) {
		r.From = v
	}
}

// WithSize - specifies a max number of filters to get.
//
func (f MlGetFilters) WithSize(v interface{}) func(*MlGetFiltersRequest) {
	return func(r *MlGetFiltersRequest) {
		r.Size = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f MlGetFilters) WithPretty() func(*MlGetFiltersRequest) {
	return func(r *MlGetFiltersRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f MlGetFilters) WithHuman() func(*MlGetFiltersRequest) {
	return func(r *MlGetFiltersRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f MlGetFilters) WithErrorTrace() func(*MlGetFiltersRequest) {
	return func(r *MlGetFiltersRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f MlGetFilters) WithFilterPath(v ...string) func(*MlGetFiltersRequest) {
	return func(r *MlGetFiltersRequest) {
		r.FilterPath = v
	}
}
