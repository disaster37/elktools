// Code generated from specification version 7.0.0-SNAPSHOT: DO NOT EDIT

package xpack

import (
	"context"
	"strings"
)

func newWatcherGetWatchFunc(t Transport) WatcherGetWatch {
	return func(id string, o ...func(*WatcherGetWatchRequest)) (*Response, error) {
		var r = WatcherGetWatchRequest{DocumentID: id}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

//
// See full documentation at http://www.elastic.co/guide/en/elasticsearch/reference/current/watcher-api-get-watch.html.
//
type WatcherGetWatch func(id string, o ...func(*WatcherGetWatchRequest)) (*Response, error)

// WatcherGetWatchRequest configures the Watcher  Get Watch API request.
//
type WatcherGetWatchRequest struct {
	DocumentID string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r WatcherGetWatchRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "GET"

	path.Grow(1 + len("_watcher") + 1 + len("watch") + 1 + len(r.DocumentID))
	path.WriteString("/")
	path.WriteString("_watcher")
	path.WriteString("/")
	path.WriteString("watch")
	path.WriteString("/")
	path.WriteString(r.DocumentID)

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
func (f WatcherGetWatch) WithContext(v context.Context) func(*WatcherGetWatchRequest) {
	return func(r *WatcherGetWatchRequest) {
		r.ctx = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f WatcherGetWatch) WithPretty() func(*WatcherGetWatchRequest) {
	return func(r *WatcherGetWatchRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f WatcherGetWatch) WithHuman() func(*WatcherGetWatchRequest) {
	return func(r *WatcherGetWatchRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f WatcherGetWatch) WithErrorTrace() func(*WatcherGetWatchRequest) {
	return func(r *WatcherGetWatchRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f WatcherGetWatch) WithFilterPath(v ...string) func(*WatcherGetWatchRequest) {
	return func(r *WatcherGetWatchRequest) {
		r.FilterPath = v
	}
}
