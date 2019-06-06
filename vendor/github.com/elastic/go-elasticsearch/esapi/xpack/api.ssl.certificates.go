// Code generated from specification version 7.0.0-SNAPSHOT: DO NOT EDIT

package xpack

import (
	"context"
	"strings"
)

func newSslCertificatesFunc(t Transport) SslCertificates {
	return func(o ...func(*SslCertificatesRequest)) (*Response, error) {
		var r = SslCertificatesRequest{}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-ssl.html.
//
type SslCertificates func(o ...func(*SslCertificatesRequest)) (*Response, error)

// SslCertificatesRequest configures the Ssl Certificates API request.
//
type SslCertificatesRequest struct {
	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r SslCertificatesRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "GET"

	path.Grow(len("/_ssl/certificates"))
	path.WriteString("/_ssl/certificates")

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
func (f SslCertificates) WithContext(v context.Context) func(*SslCertificatesRequest) {
	return func(r *SslCertificatesRequest) {
		r.ctx = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f SslCertificates) WithPretty() func(*SslCertificatesRequest) {
	return func(r *SslCertificatesRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f SslCertificates) WithHuman() func(*SslCertificatesRequest) {
	return func(r *SslCertificatesRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f SslCertificates) WithErrorTrace() func(*SslCertificatesRequest) {
	return func(r *SslCertificatesRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f SslCertificates) WithFilterPath(v ...string) func(*SslCertificatesRequest) {
	return func(r *SslCertificatesRequest) {
		r.FilterPath = v
	}
}
