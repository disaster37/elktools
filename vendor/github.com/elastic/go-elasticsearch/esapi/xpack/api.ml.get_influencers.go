// Code generated from specification version 7.0.0-SNAPSHOT: DO NOT EDIT

package xpack

import (
	"context"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func newMlGetInfluencersFunc(t Transport) MlGetInfluencers {
	return func(job_id string, o ...func(*MlGetInfluencersRequest)) (*Response, error) {
		var r = MlGetInfluencersRequest{JobID: job_id}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

//
// See full documentation at http://www.elastic.co/guide/en/elasticsearch/reference/current/ml-get-influencer.html.
//
type MlGetInfluencers func(job_id string, o ...func(*MlGetInfluencersRequest)) (*Response, error)

// MlGetInfluencersRequest configures the Ml  Get Influencers API request.
//
type MlGetInfluencersRequest struct {
	Body io.Reader

	JobID           string
	Desc            *bool
	End             string
	ExcludeInterim  *bool
	From            interface{}
	InfluencerScore interface{}
	Size            interface{}
	Sort            string
	Start           string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r MlGetInfluencersRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "GET"

	path.Grow(1 + len("_ml") + 1 + len("anomaly_detectors") + 1 + len(r.JobID) + 1 + len("results") + 1 + len("influencers"))
	path.WriteString("/")
	path.WriteString("_ml")
	path.WriteString("/")
	path.WriteString("anomaly_detectors")
	path.WriteString("/")
	path.WriteString(r.JobID)
	path.WriteString("/")
	path.WriteString("results")
	path.WriteString("/")
	path.WriteString("influencers")

	params = make(map[string]string)

	if r.Desc != nil {
		params["desc"] = strconv.FormatBool(*r.Desc)
	}

	if r.End != "" {
		params["end"] = r.End
	}

	if r.ExcludeInterim != nil {
		params["exclude_interim"] = strconv.FormatBool(*r.ExcludeInterim)
	}

	if r.From != nil {
		params["from"] = fmt.Sprintf("%v", r.From)
	}

	if r.InfluencerScore != nil {
		params["influencer_score"] = fmt.Sprintf("%v", r.InfluencerScore)
	}

	if r.Size != nil {
		params["size"] = fmt.Sprintf("%v", r.Size)
	}

	if r.Sort != "" {
		params["sort"] = r.Sort
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
func (f MlGetInfluencers) WithContext(v context.Context) func(*MlGetInfluencersRequest) {
	return func(r *MlGetInfluencersRequest) {
		r.ctx = v
	}
}

// WithBody - Influencer selection criteria.
//
func (f MlGetInfluencers) WithBody(v io.Reader) func(*MlGetInfluencersRequest) {
	return func(r *MlGetInfluencersRequest) {
		r.Body = v
	}
}

// WithDesc - whether the results should be sorted in decending order.
//
func (f MlGetInfluencers) WithDesc(v bool) func(*MlGetInfluencersRequest) {
	return func(r *MlGetInfluencersRequest) {
		r.Desc = &v
	}
}

// WithEnd - end timestamp for the requested influencers.
//
func (f MlGetInfluencers) WithEnd(v string) func(*MlGetInfluencersRequest) {
	return func(r *MlGetInfluencersRequest) {
		r.End = v
	}
}

// WithExcludeInterim - exclude interim results.
//
func (f MlGetInfluencers) WithExcludeInterim(v bool) func(*MlGetInfluencersRequest) {
	return func(r *MlGetInfluencersRequest) {
		r.ExcludeInterim = &v
	}
}

// WithFrom - skips a number of influencers.
//
func (f MlGetInfluencers) WithFrom(v interface{}) func(*MlGetInfluencersRequest) {
	return func(r *MlGetInfluencersRequest) {
		r.From = v
	}
}

// WithInfluencerScore - influencer score threshold for the requested influencers.
//
func (f MlGetInfluencers) WithInfluencerScore(v interface{}) func(*MlGetInfluencersRequest) {
	return func(r *MlGetInfluencersRequest) {
		r.InfluencerScore = v
	}
}

// WithSize - specifies a max number of influencers to get.
//
func (f MlGetInfluencers) WithSize(v interface{}) func(*MlGetInfluencersRequest) {
	return func(r *MlGetInfluencersRequest) {
		r.Size = v
	}
}

// WithSort - sort field for the requested influencers.
//
func (f MlGetInfluencers) WithSort(v string) func(*MlGetInfluencersRequest) {
	return func(r *MlGetInfluencersRequest) {
		r.Sort = v
	}
}

// WithStart - start timestamp for the requested influencers.
//
func (f MlGetInfluencers) WithStart(v string) func(*MlGetInfluencersRequest) {
	return func(r *MlGetInfluencersRequest) {
		r.Start = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f MlGetInfluencers) WithPretty() func(*MlGetInfluencersRequest) {
	return func(r *MlGetInfluencersRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f MlGetInfluencers) WithHuman() func(*MlGetInfluencersRequest) {
	return func(r *MlGetInfluencersRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f MlGetInfluencers) WithErrorTrace() func(*MlGetInfluencersRequest) {
	return func(r *MlGetInfluencersRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f MlGetInfluencers) WithFilterPath(v ...string) func(*MlGetInfluencersRequest) {
	return func(r *MlGetInfluencersRequest) {
		r.FilterPath = v
	}
}
