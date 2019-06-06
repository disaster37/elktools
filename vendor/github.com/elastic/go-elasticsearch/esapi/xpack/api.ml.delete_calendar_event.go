// Code generated from specification version 7.0.0-SNAPSHOT: DO NOT EDIT

package xpack

import (
	"context"
	"strings"
)

func newMlDeleteCalendarEventFunc(t Transport) MlDeleteCalendarEvent {
	return func(calendar_id string, event_id string, o ...func(*MlDeleteCalendarEventRequest)) (*Response, error) {
		var r = MlDeleteCalendarEventRequest{CalendarID: calendar_id, EventID: event_id}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

//
type MlDeleteCalendarEvent func(calendar_id string, event_id string, o ...func(*MlDeleteCalendarEventRequest)) (*Response, error)

// MlDeleteCalendarEventRequest configures the Ml   Delete Calendar Event API request.
//
type MlDeleteCalendarEventRequest struct {
	CalendarID string
	EventID    string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r MlDeleteCalendarEventRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "DELETE"

	path.Grow(1 + len("_ml") + 1 + len("calendars") + 1 + len(r.CalendarID) + 1 + len("events") + 1 + len(r.EventID))
	path.WriteString("/")
	path.WriteString("_ml")
	path.WriteString("/")
	path.WriteString("calendars")
	path.WriteString("/")
	path.WriteString(r.CalendarID)
	path.WriteString("/")
	path.WriteString("events")
	path.WriteString("/")
	path.WriteString(r.EventID)

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
func (f MlDeleteCalendarEvent) WithContext(v context.Context) func(*MlDeleteCalendarEventRequest) {
	return func(r *MlDeleteCalendarEventRequest) {
		r.ctx = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f MlDeleteCalendarEvent) WithPretty() func(*MlDeleteCalendarEventRequest) {
	return func(r *MlDeleteCalendarEventRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f MlDeleteCalendarEvent) WithHuman() func(*MlDeleteCalendarEventRequest) {
	return func(r *MlDeleteCalendarEventRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f MlDeleteCalendarEvent) WithErrorTrace() func(*MlDeleteCalendarEventRequest) {
	return func(r *MlDeleteCalendarEventRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f MlDeleteCalendarEvent) WithFilterPath(v ...string) func(*MlDeleteCalendarEventRequest) {
	return func(r *MlDeleteCalendarEventRequest) {
		r.FilterPath = v
	}
}
