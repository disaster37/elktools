// Code generated from specification version 7.0.0-SNAPSHOT: DO NOT EDIT

package xpack

import (
	"context"
	"io"
	"strings"
)

func newMlPutCalendarFunc(t Transport) MlPutCalendar {
	return func(calendar_id string, o ...func(*MlPutCalendarRequest)) (*Response, error) {
		var r = MlPutCalendarRequest{CalendarID: calendar_id}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

//
type MlPutCalendar func(calendar_id string, o ...func(*MlPutCalendarRequest)) (*Response, error)

// MlPutCalendarRequest configures the Ml  Put Calendar API request.
//
type MlPutCalendarRequest struct {
	Body io.Reader

	CalendarID string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r MlPutCalendarRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "PUT"

	path.Grow(1 + len("_ml") + 1 + len("calendars") + 1 + len(r.CalendarID))
	path.WriteString("/")
	path.WriteString("_ml")
	path.WriteString("/")
	path.WriteString("calendars")
	path.WriteString("/")
	path.WriteString(r.CalendarID)

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
func (f MlPutCalendar) WithContext(v context.Context) func(*MlPutCalendarRequest) {
	return func(r *MlPutCalendarRequest) {
		r.ctx = v
	}
}

// WithBody - The calendar details.
//
func (f MlPutCalendar) WithBody(v io.Reader) func(*MlPutCalendarRequest) {
	return func(r *MlPutCalendarRequest) {
		r.Body = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f MlPutCalendar) WithPretty() func(*MlPutCalendarRequest) {
	return func(r *MlPutCalendarRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f MlPutCalendar) WithHuman() func(*MlPutCalendarRequest) {
	return func(r *MlPutCalendarRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f MlPutCalendar) WithErrorTrace() func(*MlPutCalendarRequest) {
	return func(r *MlPutCalendarRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f MlPutCalendar) WithFilterPath(v ...string) func(*MlPutCalendarRequest) {
	return func(r *MlPutCalendarRequest) {
		r.FilterPath = v
	}
}
