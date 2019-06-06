// Code generated from specification version 7.0.0-SNAPSHOT: DO NOT EDIT

package xpack

import (
	"context"
	"strconv"
	"strings"
	"time"
)

func newMlDeleteForecastFunc(t Transport) MlDeleteForecast {
	return func(job_id string, o ...func(*MlDeleteForecastRequest)) (*Response, error) {
		var r = MlDeleteForecastRequest{JobID: job_id}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

//
// See full documentation at http://www.elastic.co/guide/en/elasticsearch/reference/current/ml-delete-forecast.html.
//
type MlDeleteForecast func(job_id string, o ...func(*MlDeleteForecastRequest)) (*Response, error)

// MlDeleteForecastRequest configures the Ml  Delete Forecast API request.
//
type MlDeleteForecastRequest struct {
	JobID            string
	ForecastID       string
	AllowNoForecasts *bool
	Timeout          time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r MlDeleteForecastRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "DELETE"

	path.Grow(1 + len("_ml") + 1 + len("anomaly_detectors") + 1 + len(r.JobID) + 1 + len("_forecast") + 1 + len(r.ForecastID))
	path.WriteString("/")
	path.WriteString("_ml")
	path.WriteString("/")
	path.WriteString("anomaly_detectors")
	path.WriteString("/")
	path.WriteString(r.JobID)
	path.WriteString("/")
	path.WriteString("_forecast")
	if r.ForecastID != "" {
		path.WriteString("/")
		path.WriteString(r.ForecastID)
	}

	params = make(map[string]string)

	if r.AllowNoForecasts != nil {
		params["allow_no_forecasts"] = strconv.FormatBool(*r.AllowNoForecasts)
	}

	if r.Timeout != 0 {
		params["timeout"] = formatDuration(r.Timeout)
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
func (f MlDeleteForecast) WithContext(v context.Context) func(*MlDeleteForecastRequest) {
	return func(r *MlDeleteForecastRequest) {
		r.ctx = v
	}
}

// WithForecastID - the ID of the forecast to delete, can be comma delimited list. leaving blank implies `_all`.
//
func (f MlDeleteForecast) WithForecastID(v string) func(*MlDeleteForecastRequest) {
	return func(r *MlDeleteForecastRequest) {
		r.ForecastID = v
	}
}

// WithAllowNoForecasts - whether to ignore if `_all` matches no forecasts.
//
func (f MlDeleteForecast) WithAllowNoForecasts(v bool) func(*MlDeleteForecastRequest) {
	return func(r *MlDeleteForecastRequest) {
		r.AllowNoForecasts = &v
	}
}

// WithTimeout - controls the time to wait until the forecast(s) are deleted. default to 30 seconds.
//
func (f MlDeleteForecast) WithTimeout(v time.Duration) func(*MlDeleteForecastRequest) {
	return func(r *MlDeleteForecastRequest) {
		r.Timeout = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f MlDeleteForecast) WithPretty() func(*MlDeleteForecastRequest) {
	return func(r *MlDeleteForecastRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f MlDeleteForecast) WithHuman() func(*MlDeleteForecastRequest) {
	return func(r *MlDeleteForecastRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f MlDeleteForecast) WithErrorTrace() func(*MlDeleteForecastRequest) {
	return func(r *MlDeleteForecastRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f MlDeleteForecast) WithFilterPath(v ...string) func(*MlDeleteForecastRequest) {
	return func(r *MlDeleteForecastRequest) {
		r.FilterPath = v
	}
}
