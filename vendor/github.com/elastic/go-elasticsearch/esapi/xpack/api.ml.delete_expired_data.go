// Code generated from specification version 7.0.0-SNAPSHOT: DO NOT EDIT

package xpack

import (
	"context"
	"strings"
)

func newMlDeleteExpiredDataFunc(t Transport) MlDeleteExpiredData {
	return func(o ...func(*MlDeleteExpiredDataRequest)) (*Response, error) {
		var r = MlDeleteExpiredDataRequest{}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

//
type MlDeleteExpiredData func(o ...func(*MlDeleteExpiredDataRequest)) (*Response, error)

// MlDeleteExpiredDataRequest configures the Ml   Delete Expired Data API request.
//
type MlDeleteExpiredDataRequest struct {
	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r MlDeleteExpiredDataRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "DELETE"

	path.Grow(len("/_ml/_delete_expired_data"))
	path.WriteString("/_ml/_delete_expired_data")

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
func (f MlDeleteExpiredData) WithContext(v context.Context) func(*MlDeleteExpiredDataRequest) {
	return func(r *MlDeleteExpiredDataRequest) {
		r.ctx = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f MlDeleteExpiredData) WithPretty() func(*MlDeleteExpiredDataRequest) {
	return func(r *MlDeleteExpiredDataRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f MlDeleteExpiredData) WithHuman() func(*MlDeleteExpiredDataRequest) {
	return func(r *MlDeleteExpiredDataRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f MlDeleteExpiredData) WithErrorTrace() func(*MlDeleteExpiredDataRequest) {
	return func(r *MlDeleteExpiredDataRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f MlDeleteExpiredData) WithFilterPath(v ...string) func(*MlDeleteExpiredDataRequest) {
	return func(r *MlDeleteExpiredDataRequest) {
		r.FilterPath = v
	}
}
