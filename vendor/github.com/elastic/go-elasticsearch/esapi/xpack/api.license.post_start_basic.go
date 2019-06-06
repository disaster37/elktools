// Code generated from specification version 7.0.0-SNAPSHOT: DO NOT EDIT

package xpack

import (
	"context"
	"strconv"
	"strings"
)

func newLicensePostStartBasicFunc(t Transport) LicensePostStartBasic {
	return func(o ...func(*LicensePostStartBasicRequest)) (*Response, error) {
		var r = LicensePostStartBasicRequest{}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

//
// See full documentation at https://www.elastic.co/guide/en/x-pack/current/license-management.html.
//
type LicensePostStartBasic func(o ...func(*LicensePostStartBasicRequest)) (*Response, error)

// LicensePostStartBasicRequest configures the License   Post Start Basic API request.
//
type LicensePostStartBasicRequest struct {
	Acknowledge *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r LicensePostStartBasicRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "POST"

	path.Grow(len("/_license/start_basic"))
	path.WriteString("/_license/start_basic")

	params = make(map[string]string)

	if r.Acknowledge != nil {
		params["acknowledge"] = strconv.FormatBool(*r.Acknowledge)
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
func (f LicensePostStartBasic) WithContext(v context.Context) func(*LicensePostStartBasicRequest) {
	return func(r *LicensePostStartBasicRequest) {
		r.ctx = v
	}
}

// WithAcknowledge - whether the user has acknowledged acknowledge messages (default: false).
//
func (f LicensePostStartBasic) WithAcknowledge(v bool) func(*LicensePostStartBasicRequest) {
	return func(r *LicensePostStartBasicRequest) {
		r.Acknowledge = &v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f LicensePostStartBasic) WithPretty() func(*LicensePostStartBasicRequest) {
	return func(r *LicensePostStartBasicRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f LicensePostStartBasic) WithHuman() func(*LicensePostStartBasicRequest) {
	return func(r *LicensePostStartBasicRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f LicensePostStartBasic) WithErrorTrace() func(*LicensePostStartBasicRequest) {
	return func(r *LicensePostStartBasicRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f LicensePostStartBasic) WithFilterPath(v ...string) func(*LicensePostStartBasicRequest) {
	return func(r *LicensePostStartBasicRequest) {
		r.FilterPath = v
	}
}
