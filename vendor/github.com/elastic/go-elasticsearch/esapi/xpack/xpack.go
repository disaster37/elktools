package xpack

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/esapi"
)

const (
	headerContentType = "Content-Type"
)

var (
	headerContentTypeJSON = []string{"application/json"}
)

type Response esapi.Response
type Transport esapi.Transport

func newRequest(method, path string, body io.Reader) (*http.Request, error) {
	r := http.Request{
		Method:     method,
		URL:        &url.URL{Path: path},
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
		Header:     make(http.Header),
	}

	if body != nil {
		switch b := body.(type) {
		case *bytes.Buffer:
			r.Body = ioutil.NopCloser(body)
			r.ContentLength = int64(b.Len())
		case *bytes.Reader:
			r.Body = ioutil.NopCloser(body)
			r.ContentLength = int64(b.Len())
		case *strings.Reader:
			r.Body = ioutil.NopCloser(body)
			r.ContentLength = int64(b.Len())
		}
	}

	return &r, nil
}

func BoolPtr(v bool) *bool { return &v }

func IntPtr(v int) *int { return &v }

func formatDuration(d time.Duration) string {
	if d < time.Millisecond {
		return strconv.FormatInt(int64(d), 10) + "nanos"
	}
	return strconv.FormatInt(int64(d)/int64(time.Millisecond), 10) + "ms"
}

func (r *Response) String() string {
	var (
		out = new(bytes.Buffer)
		b1  = bytes.NewBuffer([]byte{})
		b2  = bytes.NewBuffer([]byte{})
		tr  = io.TeeReader(r.Body, b1)
	)

	defer r.Body.Close()

	if _, err := io.Copy(b2, tr); err != nil {
		out.WriteString(fmt.Sprintf("<error reading response body: %v>", err))
		return out.String()
	}
	defer func() { r.Body = ioutil.NopCloser(b1) }()

	out.WriteString(fmt.Sprintf("[%d %s]", r.StatusCode, http.StatusText(r.StatusCode)))
	out.WriteRune(' ')
	out.ReadFrom(b2) // errcheck exclude (*bytes.Buffer).ReadFrom

	return out.String()
}

func (r *Response) Status() string {
	var b strings.Builder
	b.WriteString(strconv.Itoa(r.StatusCode))
	b.WriteString(" ")
	b.WriteString(http.StatusText(r.StatusCode))
	return b.String()
}

func (r *Response) IsError() bool {
	return r.StatusCode > 299
}
