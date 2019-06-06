// Code generated from specification version 7.0.0-SNAPSHOT: DO NOT EDIT

package xpack

import (
	"context"
	"fmt"
	"io"
	"strings"
)

func newMlGetCalendarsFunc(t Transport) MlGetCalendars {
	return func(o ...func(*MlGetCalendarsRequest)) (*Response, error) {
		var r = MlGetCalendarsRequest{}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

//
type MlGetCalendars func(o ...func(*MlGetCalendarsRequest)) (*Response, error)

// MlGetCalendarsRequest configures the Ml  Get Calendars API request.
//
type MlGetCalendarsRequest struct {
	Body io.Reader

	CalendarID string
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
func (r MlGetCalendarsRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "GET"

	path.Grow(1 + len("_ml") + 1 + len("calendars") + 1 + len(r.CalendarID))
	path.WriteString("/")
	path.WriteString("_ml")
	path.WriteString("/")
	path.WriteString("calendars")
	if r.CalendarID != "" {
		path.WriteString("/")
		path.WriteString(r.CalendarID)
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
func (f MlGetCalendars) WithContext(v context.Context) func(*MlGetCalendarsRequest) {
	return func(r *MlGetCalendarsRequest) {
		r.ctx = v
	}
}

// WithCalendarID - the ID of the calendar to fetch.
//
func (f MlGetCalendars) WithCalendarID(v string) func(*MlGetCalendarsRequest) {
	return func(r *MlGetCalendarsRequest) {
		r.CalendarID = v
	}
}

// WithBody - The from and size parameters optionally sent in the body.
//
func (f MlGetCalendars) WithBody(v io.Reader) func(*MlGetCalendarsRequest) {
	return func(r *MlGetCalendarsRequest) {
		r.Body = v
	}
}

// WithFrom - skips a number of calendars.
//
func (f MlGetCalendars) WithFrom(v interface{}) func(*MlGetCalendarsRequest) {
	return func(r *MlGetCalendarsRequest) {
		r.From = v
	}
}

// WithSize - specifies a max number of calendars to get.
//
func (f MlGetCalendars) WithSize(v interface{}) func(*MlGetCalendarsRequest) {
	return func(r *MlGetCalendarsRequest) {
		r.Size = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f MlGetCalendars) WithPretty() func(*MlGetCalendarsRequest) {
	return func(r *MlGetCalendarsRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f MlGetCalendars) WithHuman() func(*MlGetCalendarsRequest) {
	return func(r *MlGetCalendarsRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f MlGetCalendars) WithErrorTrace() func(*MlGetCalendarsRequest) {
	return func(r *MlGetCalendarsRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f MlGetCalendars) WithFilterPath(v ...string) func(*MlGetCalendarsRequest) {
	return func(r *MlGetCalendarsRequest) {
		r.FilterPath = v
	}
}
