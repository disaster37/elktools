// Code generated from specification version 7.0.0-SNAPSHOT: DO NOT EDIT

package xpack

import (
	"context"
	"strconv"
	"strings"
	"time"
)

func newMlSetUpgradeModeFunc(t Transport) MlSetUpgradeMode {
	return func(o ...func(*MlSetUpgradeModeRequest)) (*Response, error) {
		var r = MlSetUpgradeModeRequest{}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

//
// See full documentation at http://www.elastic.co/guide/en/elasticsearch/reference/current/ml-set-upgrade-mode.html.
//
type MlSetUpgradeMode func(o ...func(*MlSetUpgradeModeRequest)) (*Response, error)

// MlSetUpgradeModeRequest configures the Ml   Set Upgrade Mode API request.
//
type MlSetUpgradeModeRequest struct {
	Enabled *bool
	Timeout time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r MlSetUpgradeModeRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "POST"

	path.Grow(len("/_ml/set_upgrade_mode"))
	path.WriteString("/_ml/set_upgrade_mode")

	params = make(map[string]string)

	if r.Enabled != nil {
		params["enabled"] = strconv.FormatBool(*r.Enabled)
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
func (f MlSetUpgradeMode) WithContext(v context.Context) func(*MlSetUpgradeModeRequest) {
	return func(r *MlSetUpgradeModeRequest) {
		r.ctx = v
	}
}

// WithEnabled - whether to enable upgrade_mode ml setting or not. defaults to false..
//
func (f MlSetUpgradeMode) WithEnabled(v bool) func(*MlSetUpgradeModeRequest) {
	return func(r *MlSetUpgradeModeRequest) {
		r.Enabled = &v
	}
}

// WithTimeout - controls the time to wait before action times out. defaults to 30 seconds.
//
func (f MlSetUpgradeMode) WithTimeout(v time.Duration) func(*MlSetUpgradeModeRequest) {
	return func(r *MlSetUpgradeModeRequest) {
		r.Timeout = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f MlSetUpgradeMode) WithPretty() func(*MlSetUpgradeModeRequest) {
	return func(r *MlSetUpgradeModeRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f MlSetUpgradeMode) WithHuman() func(*MlSetUpgradeModeRequest) {
	return func(r *MlSetUpgradeModeRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f MlSetUpgradeMode) WithErrorTrace() func(*MlSetUpgradeModeRequest) {
	return func(r *MlSetUpgradeModeRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f MlSetUpgradeMode) WithFilterPath(v ...string) func(*MlSetUpgradeModeRequest) {
	return func(r *MlSetUpgradeModeRequest) {
		r.FilterPath = v
	}
}
