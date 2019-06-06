// Code generated from specification version 7.0.0-SNAPSHOT: DO NOT EDIT

package xpack

import (
	"context"
	"strings"
)

func newMlPutCalendarJobFunc(t Transport) MlPutCalendarJob {
	return func(calendar_id string, job_id string, o ...func(*MlPutCalendarJobRequest)) (*Response, error) {
		var r = MlPutCalendarJobRequest{CalendarID: calendar_id, JobID: job_id}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

//
type MlPutCalendarJob func(calendar_id string, job_id string, o ...func(*MlPutCalendarJobRequest)) (*Response, error)

// MlPutCalendarJobRequest configures the Ml   Put Calendar Job API request.
//
type MlPutCalendarJobRequest struct {
	CalendarID string
	JobID      string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r MlPutCalendarJobRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "PUT"

	path.Grow(1 + len("_ml") + 1 + len("calendars") + 1 + len(r.CalendarID) + 1 + len("jobs") + 1 + len(r.JobID))
	path.WriteString("/")
	path.WriteString("_ml")
	path.WriteString("/")
	path.WriteString("calendars")
	path.WriteString("/")
	path.WriteString(r.CalendarID)
	path.WriteString("/")
	path.WriteString("jobs")
	path.WriteString("/")
	path.WriteString(r.JobID)

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
func (f MlPutCalendarJob) WithContext(v context.Context) func(*MlPutCalendarJobRequest) {
	return func(r *MlPutCalendarJobRequest) {
		r.ctx = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f MlPutCalendarJob) WithPretty() func(*MlPutCalendarJobRequest) {
	return func(r *MlPutCalendarJobRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f MlPutCalendarJob) WithHuman() func(*MlPutCalendarJobRequest) {
	return func(r *MlPutCalendarJobRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f MlPutCalendarJob) WithErrorTrace() func(*MlPutCalendarJobRequest) {
	return func(r *MlPutCalendarJobRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f MlPutCalendarJob) WithFilterPath(v ...string) func(*MlPutCalendarJobRequest) {
	return func(r *MlPutCalendarJobRequest) {
		r.FilterPath = v
	}
}
