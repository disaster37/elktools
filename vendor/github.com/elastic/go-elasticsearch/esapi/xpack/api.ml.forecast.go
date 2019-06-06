// Code generated from specification version 7.0.0-SNAPSHOT: DO NOT EDIT

package xpack

import (
	"context"
	"strings"
	"time"
)

func newMlForecastFunc(t Transport) MlForecast {
	return func(job_id string, o ...func(*MlForecastRequest)) (*Response, error) {
		var r = MlForecastRequest{JobID: job_id}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

//
type MlForecast func(job_id string, o ...func(*MlForecastRequest)) (*Response, error)

// MlForecastRequest configures the Ml Forecast API request.
//
type MlForecastRequest struct {
	JobID     string
	Duration  time.Duration
	ExpiresIn time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r MlForecastRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "POST"

	path.Grow(1 + len("_ml") + 1 + len("anomaly_detectors") + 1 + len(r.JobID) + 1 + len("_forecast"))
	path.WriteString("/")
	path.WriteString("_ml")
	path.WriteString("/")
	path.WriteString("anomaly_detectors")
	path.WriteString("/")
	path.WriteString(r.JobID)
	path.WriteString("/")
	path.WriteString("_forecast")

	params = make(map[string]string)

	if r.Duration != 0 {
		params["duration"] = formatDuration(r.Duration)
	}

	if r.ExpiresIn != 0 {
		params["expires_in"] = formatDuration(r.ExpiresIn)
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
func (f MlForecast) WithContext(v context.Context) func(*MlForecastRequest) {
	return func(r *MlForecastRequest) {
		r.ctx = v
	}
}

// WithDuration - the duration of the forecast.
//
func (f MlForecast) WithDuration(v time.Duration) func(*MlForecastRequest) {
	return func(r *MlForecastRequest) {
		r.Duration = v
	}
}

// WithExpiresIn - the time interval after which the forecast expires. expired forecasts will be deleted at the first opportunity..
//
func (f MlForecast) WithExpiresIn(v time.Duration) func(*MlForecastRequest) {
	return func(r *MlForecastRequest) {
		r.ExpiresIn = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f MlForecast) WithPretty() func(*MlForecastRequest) {
	return func(r *MlForecastRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f MlForecast) WithHuman() func(*MlForecastRequest) {
	return func(r *MlForecastRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f MlForecast) WithErrorTrace() func(*MlForecastRequest) {
	return func(r *MlForecastRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f MlForecast) WithFilterPath(v ...string) func(*MlForecastRequest) {
	return func(r *MlForecastRequest) {
		r.FilterPath = v
	}
}
