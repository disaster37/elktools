// Code generated from specification version 7.0.0-SNAPSHOT: DO NOT EDIT

package xpack

import (
	"context"
	"strconv"
	"strings"
)

func newLicenseGetFunc(t Transport) LicenseGet {
	return func(o ...func(*LicenseGetRequest)) (*Response, error) {
		var r = LicenseGetRequest{}
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
type LicenseGet func(o ...func(*LicenseGetRequest)) (*Response, error)

// LicenseGetRequest configures the License Get API request.
//
type LicenseGetRequest struct {
	Local *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r LicenseGetRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "GET"

	path.Grow(len("/_license"))
	path.WriteString("/_license")

	params = make(map[string]string)

	if r.Local != nil {
		params["local"] = strconv.FormatBool(*r.Local)
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
func (f LicenseGet) WithContext(v context.Context) func(*LicenseGetRequest) {
	return func(r *LicenseGetRequest) {
		r.ctx = v
	}
}

// WithLocal - return local information, do not retrieve the state from master node (default: false).
//
func (f LicenseGet) WithLocal(v bool) func(*LicenseGetRequest) {
	return func(r *LicenseGetRequest) {
		r.Local = &v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f LicenseGet) WithPretty() func(*LicenseGetRequest) {
	return func(r *LicenseGetRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f LicenseGet) WithHuman() func(*LicenseGetRequest) {
	return func(r *LicenseGetRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f LicenseGet) WithErrorTrace() func(*LicenseGetRequest) {
	return func(r *LicenseGetRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f LicenseGet) WithFilterPath(v ...string) func(*LicenseGetRequest) {
	return func(r *LicenseGetRequest) {
		r.FilterPath = v
	}
}
