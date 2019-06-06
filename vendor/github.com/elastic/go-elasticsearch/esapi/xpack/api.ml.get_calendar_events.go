// Code generated from specification version 7.0.0-SNAPSHOT: DO NOT EDIT

package xpack

import (
	"context"
	"fmt"
	"strings"
)

func newMlGetCalendarEventsFunc(t Transport) MlGetCalendarEvents {
	return func(calendar_id string, o ...func(*MlGetCalendarEventsRequest)) (*Response, error) {
		var r = MlGetCalendarEventsRequest{CalendarID: calendar_id}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

//
type MlGetCalendarEvents func(calendar_id string, o ...func(*MlGetCalendarEventsRequest)) (*Response, error)

// MlGetCalendarEventsRequest configures the Ml   Get Calendar Events API request.
//
type MlGetCalendarEventsRequest struct {
	CalendarID string
	End        interface{}
	From       interface{}
	JobID      string
	Size       interface{}
	Start      string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r MlGetCalendarEventsRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "GET"

	path.Grow(1 + len("_ml") + 1 + len("calendars") + 1 + len(r.CalendarID) + 1 + len("events"))
	path.WriteString("/")
	path.WriteString("_ml")
	path.WriteString("/")
	path.WriteString("calendars")
	path.WriteString("/")
	path.WriteString(r.CalendarID)
	path.WriteString("/")
	path.WriteString("events")

	params = make(map[string]string)

	if r.End != nil {
		params["end"] = fmt.Sprintf("%v", r.End)
	}

	if r.From != nil {
		params["from"] = fmt.Sprintf("%v", r.From)
	}

	if r.JobID != "" {
		params["job_id"] = r.JobID
	}

	if r.Size != nil {
		params["size"] = fmt.Sprintf("%v", r.Size)
	}

	if r.Start != "" {
		params["start"] = r.Start
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
func (f MlGetCalendarEvents) WithContext(v context.Context) func(*MlGetCalendarEventsRequest) {
	return func(r *MlGetCalendarEventsRequest) {
		r.ctx = v
	}
}

// WithEnd - get events before this time.
//
func (f MlGetCalendarEvents) WithEnd(v interface{}) func(*MlGetCalendarEventsRequest) {
	return func(r *MlGetCalendarEventsRequest) {
		r.End = v
	}
}

// WithFrom - skips a number of events.
//
func (f MlGetCalendarEvents) WithFrom(v interface{}) func(*MlGetCalendarEventsRequest) {
	return func(r *MlGetCalendarEventsRequest) {
		r.From = v
	}
}

// WithJobID - get events for the job. when this option is used calendar_id must be '_all'.
//
func (f MlGetCalendarEvents) WithJobID(v string) func(*MlGetCalendarEventsRequest) {
	return func(r *MlGetCalendarEventsRequest) {
		r.JobID = v
	}
}

// WithSize - specifies a max number of events to get.
//
func (f MlGetCalendarEvents) WithSize(v interface{}) func(*MlGetCalendarEventsRequest) {
	return func(r *MlGetCalendarEventsRequest) {
		r.Size = v
	}
}

// WithStart - get events after this time.
//
func (f MlGetCalendarEvents) WithStart(v string) func(*MlGetCalendarEventsRequest) {
	return func(r *MlGetCalendarEventsRequest) {
		r.Start = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f MlGetCalendarEvents) WithPretty() func(*MlGetCalendarEventsRequest) {
	return func(r *MlGetCalendarEventsRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f MlGetCalendarEvents) WithHuman() func(*MlGetCalendarEventsRequest) {
	return func(r *MlGetCalendarEventsRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f MlGetCalendarEvents) WithErrorTrace() func(*MlGetCalendarEventsRequest) {
	return func(r *MlGetCalendarEventsRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f MlGetCalendarEvents) WithFilterPath(v ...string) func(*MlGetCalendarEventsRequest) {
	return func(r *MlGetCalendarEventsRequest) {
		r.FilterPath = v
	}
}
