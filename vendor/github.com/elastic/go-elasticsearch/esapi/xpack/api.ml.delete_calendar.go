// Code generated from specification version 7.0.0-SNAPSHOT: DO NOT EDIT

package xpack

import (
	"context"
	"strings"
)

func newMlDeleteCalendarFunc(t Transport) MlDeleteCalendar {
	return func(calendar_id string, o ...func(*MlDeleteCalendarRequest)) (*Response, error) {
		var r = MlDeleteCalendarRequest{CalendarID: calendar_id}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

//
type MlDeleteCalendar func(calendar_id string, o ...func(*MlDeleteCalendarRequest)) (*Response, error)

// MlDeleteCalendarRequest configures the Ml  Delete Calendar API request.
//
type MlDeleteCalendarRequest struct {
	CalendarID string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r MlDeleteCalendarRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "DELETE"

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
func (f MlDeleteCalendar) WithContext(v context.Context) func(*MlDeleteCalendarRequest) {
	return func(r *MlDeleteCalendarRequest) {
		r.ctx = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f MlDeleteCalendar) WithPretty() func(*MlDeleteCalendarRequest) {
	return func(r *MlDeleteCalendarRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f MlDeleteCalendar) WithHuman() func(*MlDeleteCalendarRequest) {
	return func(r *MlDeleteCalendarRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f MlDeleteCalendar) WithErrorTrace() func(*MlDeleteCalendarRequest) {
	return func(r *MlDeleteCalendarRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f MlDeleteCalendar) WithFilterPath(v ...string) func(*MlDeleteCalendarRequest) {
	return func(r *MlDeleteCalendarRequest) {
		r.FilterPath = v
	}
}