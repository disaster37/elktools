// Code generated from specification version 7.0.0-SNAPSHOT: DO NOT EDIT

package xpack

import (
	"context"
	"strings"
)

func newSecurityGetApiKeyFunc(t Transport) SecurityGetApiKey {
	return func(o ...func(*SecurityGetApiKeyRequest)) (*Response, error) {
		var r = SecurityGetApiKeyRequest{}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-get-api-key.html.
//
type SecurityGetApiKey func(o ...func(*SecurityGetApiKeyRequest)) (*Response, error)

// SecurityGetApiKeyRequest configures the Security   Get Api Key API request.
//
type SecurityGetApiKeyRequest struct {
	DocumentID string
	Name       string
	RealmName  string
	Username   string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r SecurityGetApiKeyRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "GET"

	path.Grow(len("/_security/api_key"))
	path.WriteString("/_security/api_key")

	params = make(map[string]string)

	if r.DocumentID != "" {
		params["id"] = r.DocumentID
	}

	if r.Name != "" {
		params["name"] = r.Name
	}

	if r.RealmName != "" {
		params["realm_name"] = r.RealmName
	}

	if r.Username != "" {
		params["username"] = r.Username
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
func (f SecurityGetApiKey) WithContext(v context.Context) func(*SecurityGetApiKeyRequest) {
	return func(r *SecurityGetApiKeyRequest) {
		r.ctx = v
	}
}

// WithDocumentID - api key ID of the api key to be retrieved.
//
func (f SecurityGetApiKey) WithDocumentID(v string) func(*SecurityGetApiKeyRequest) {
	return func(r *SecurityGetApiKeyRequest) {
		r.DocumentID = v
	}
}

// WithName - api key name of the api key to be retrieved.
//
func (f SecurityGetApiKey) WithName(v string) func(*SecurityGetApiKeyRequest) {
	return func(r *SecurityGetApiKeyRequest) {
		r.Name = v
	}
}

// WithRealmName - realm name of the user who created this api key to be retrieved.
//
func (f SecurityGetApiKey) WithRealmName(v string) func(*SecurityGetApiKeyRequest) {
	return func(r *SecurityGetApiKeyRequest) {
		r.RealmName = v
	}
}

// WithUsername - user name of the user who created this api key to be retrieved.
//
func (f SecurityGetApiKey) WithUsername(v string) func(*SecurityGetApiKeyRequest) {
	return func(r *SecurityGetApiKeyRequest) {
		r.Username = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f SecurityGetApiKey) WithPretty() func(*SecurityGetApiKeyRequest) {
	return func(r *SecurityGetApiKeyRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f SecurityGetApiKey) WithHuman() func(*SecurityGetApiKeyRequest) {
	return func(r *SecurityGetApiKeyRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f SecurityGetApiKey) WithErrorTrace() func(*SecurityGetApiKeyRequest) {
	return func(r *SecurityGetApiKeyRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f SecurityGetApiKey) WithFilterPath(v ...string) func(*SecurityGetApiKeyRequest) {
	return func(r *SecurityGetApiKeyRequest) {
		r.FilterPath = v
	}
}
