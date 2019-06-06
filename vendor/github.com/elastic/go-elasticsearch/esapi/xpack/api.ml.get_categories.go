// Code generated from specification version 7.0.0-SNAPSHOT: DO NOT EDIT

package xpack

import (
	"context"
	"fmt"
	"io"
	"strings"
)

func newMlGetCategoriesFunc(t Transport) MlGetCategories {
	return func(job_id string, o ...func(*MlGetCategoriesRequest)) (*Response, error) {
		var r = MlGetCategoriesRequest{JobID: job_id}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

//
// See full documentation at http://www.elastic.co/guide/en/elasticsearch/reference/current/ml-get-category.html.
//
type MlGetCategories func(job_id string, o ...func(*MlGetCategoriesRequest)) (*Response, error)

// MlGetCategoriesRequest configures the Ml  Get Categories API request.
//
type MlGetCategoriesRequest struct {
	Body io.Reader

	JobID      string
	CategoryID interface{}
	From       interface{}
	Size       interface{}

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r MlGetCategoriesRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "GET"

	path.Grow(1 + len("_ml") + 1 + len("anomaly_detectors") + 1 + len(r.JobID) + 1 + len("results") + 1 + len("categories"))
	path.WriteString("/")
	path.WriteString("_ml")
	path.WriteString("/")
	path.WriteString("anomaly_detectors")
	path.WriteString("/")
	path.WriteString(r.JobID)
	path.WriteString("/")
	path.WriteString("results")
	path.WriteString("/")
	path.WriteString("categories")
	// TODO: type int or long

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
func (f MlGetCategories) WithContext(v context.Context) func(*MlGetCategoriesRequest) {
	return func(r *MlGetCategoriesRequest) {
		r.ctx = v
	}
}

// WithCategoryID - the identifier of the category definition of interest.
//
func (f MlGetCategories) WithCategoryID(v interface{}) func(*MlGetCategoriesRequest) {
	return func(r *MlGetCategoriesRequest) {
		r.CategoryID = v
	}
}

// WithBody - Category selection details if not provided in URI.
//
func (f MlGetCategories) WithBody(v io.Reader) func(*MlGetCategoriesRequest) {
	return func(r *MlGetCategoriesRequest) {
		r.Body = v
	}
}

// WithFrom - skips a number of categories.
//
func (f MlGetCategories) WithFrom(v interface{}) func(*MlGetCategoriesRequest) {
	return func(r *MlGetCategoriesRequest) {
		r.From = v
	}
}

// WithSize - specifies a max number of categories to get.
//
func (f MlGetCategories) WithSize(v interface{}) func(*MlGetCategoriesRequest) {
	return func(r *MlGetCategoriesRequest) {
		r.Size = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f MlGetCategories) WithPretty() func(*MlGetCategoriesRequest) {
	return func(r *MlGetCategoriesRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f MlGetCategories) WithHuman() func(*MlGetCategoriesRequest) {
	return func(r *MlGetCategoriesRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f MlGetCategories) WithErrorTrace() func(*MlGetCategoriesRequest) {
	return func(r *MlGetCategoriesRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f MlGetCategories) WithFilterPath(v ...string) func(*MlGetCategoriesRequest) {
	return func(r *MlGetCategoriesRequest) {
		r.FilterPath = v
	}
}
