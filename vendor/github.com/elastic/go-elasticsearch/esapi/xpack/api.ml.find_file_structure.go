// Code generated from specification version 7.0.0-SNAPSHOT: DO NOT EDIT

package xpack

import (
	"context"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"
)

func newMlFindFileStructureFunc(t Transport) MlFindFileStructure {
	return func(body io.Reader, o ...func(*MlFindFileStructureRequest)) (*Response, error) {
		var r = MlFindFileStructureRequest{Body: body}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

//
// See full documentation at http://www.elastic.co/guide/en/elasticsearch/reference/current/ml-file-structure.html.
//
type MlFindFileStructure func(body io.Reader, o ...func(*MlFindFileStructureRequest)) (*Response, error)

// MlFindFileStructureRequest configures the Ml   Find File Structure API request.
//
type MlFindFileStructureRequest struct {
	Body io.Reader

	Charset          string
	ColumnNames      []string
	Delimiter        string
	Explain          *bool
	Format           string
	GrokPattern      string
	HasHeaderRow     *bool
	LinesToSample    interface{}
	Quote            string
	ShouldTrimFields *bool
	Timeout          time.Duration
	TimestampField   string
	TimestampFormat  string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r MlFindFileStructureRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "POST"

	path.Grow(len("/_ml/find_file_structure"))
	path.WriteString("/_ml/find_file_structure")

	params = make(map[string]string)

	if r.Charset != "" {
		params["charset"] = r.Charset
	}

	if len(r.ColumnNames) > 0 {
		params["column_names"] = strings.Join(r.ColumnNames, ",")
	}

	if r.Delimiter != "" {
		params["delimiter"] = r.Delimiter
	}

	if r.Explain != nil {
		params["explain"] = strconv.FormatBool(*r.Explain)
	}

	if r.Format != "" {
		params["format"] = r.Format
	}

	if r.GrokPattern != "" {
		params["grok_pattern"] = r.GrokPattern
	}

	if r.HasHeaderRow != nil {
		params["has_header_row"] = strconv.FormatBool(*r.HasHeaderRow)
	}

	if r.LinesToSample != nil {
		params["lines_to_sample"] = fmt.Sprintf("%v", r.LinesToSample)
	}

	if r.Quote != "" {
		params["quote"] = r.Quote
	}

	if r.ShouldTrimFields != nil {
		params["should_trim_fields"] = strconv.FormatBool(*r.ShouldTrimFields)
	}

	if r.Timeout != 0 {
		params["timeout"] = formatDuration(r.Timeout)
	}

	if r.TimestampField != "" {
		params["timestamp_field"] = r.TimestampField
	}

	if r.TimestampFormat != "" {
		params["timestamp_format"] = r.TimestampFormat
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
func (f MlFindFileStructure) WithContext(v context.Context) func(*MlFindFileStructureRequest) {
	return func(r *MlFindFileStructureRequest) {
		r.ctx = v
	}
}

// WithCharset - optional parameter to specify the character set of the file.
//
func (f MlFindFileStructure) WithCharset(v string) func(*MlFindFileStructureRequest) {
	return func(r *MlFindFileStructureRequest) {
		r.Charset = v
	}
}

// WithColumnNames - optional parameter containing a comma separated list of the column names for a delimited file.
//
func (f MlFindFileStructure) WithColumnNames(v ...string) func(*MlFindFileStructureRequest) {
	return func(r *MlFindFileStructureRequest) {
		r.ColumnNames = v
	}
}

// WithDelimiter - optional parameter to specify the delimiter character for a delimited file - must be a single character.
//
func (f MlFindFileStructure) WithDelimiter(v string) func(*MlFindFileStructureRequest) {
	return func(r *MlFindFileStructureRequest) {
		r.Delimiter = v
	}
}

// WithExplain - whether to include a commentary on how the structure was derived.
//
func (f MlFindFileStructure) WithExplain(v bool) func(*MlFindFileStructureRequest) {
	return func(r *MlFindFileStructureRequest) {
		r.Explain = &v
	}
}

// WithFormat - optional parameter to specify the high level file format.
//
func (f MlFindFileStructure) WithFormat(v string) func(*MlFindFileStructureRequest) {
	return func(r *MlFindFileStructureRequest) {
		r.Format = v
	}
}

// WithGrokPattern - optional parameter to specify the grok pattern that should be used to extract fields from messages in a semi-structured text file.
//
func (f MlFindFileStructure) WithGrokPattern(v string) func(*MlFindFileStructureRequest) {
	return func(r *MlFindFileStructureRequest) {
		r.GrokPattern = v
	}
}

// WithHasHeaderRow - optional parameter to specify whether a delimited file includes the column names in its first row.
//
func (f MlFindFileStructure) WithHasHeaderRow(v bool) func(*MlFindFileStructureRequest) {
	return func(r *MlFindFileStructureRequest) {
		r.HasHeaderRow = &v
	}
}

// WithLinesToSample - how many lines of the file should be included in the analysis.
//
func (f MlFindFileStructure) WithLinesToSample(v interface{}) func(*MlFindFileStructureRequest) {
	return func(r *MlFindFileStructureRequest) {
		r.LinesToSample = v
	}
}

// WithQuote - optional parameter to specify the quote character for a delimited file - must be a single character.
//
func (f MlFindFileStructure) WithQuote(v string) func(*MlFindFileStructureRequest) {
	return func(r *MlFindFileStructureRequest) {
		r.Quote = v
	}
}

// WithShouldTrimFields - optional parameter to specify whether the values between delimiters in a delimited file should have whitespace trimmed from them.
//
func (f MlFindFileStructure) WithShouldTrimFields(v bool) func(*MlFindFileStructureRequest) {
	return func(r *MlFindFileStructureRequest) {
		r.ShouldTrimFields = &v
	}
}

// WithTimeout - timeout after which the analysis will be aborted.
//
func (f MlFindFileStructure) WithTimeout(v time.Duration) func(*MlFindFileStructureRequest) {
	return func(r *MlFindFileStructureRequest) {
		r.Timeout = v
	}
}

// WithTimestampField - optional parameter to specify the timestamp field in the file.
//
func (f MlFindFileStructure) WithTimestampField(v string) func(*MlFindFileStructureRequest) {
	return func(r *MlFindFileStructureRequest) {
		r.TimestampField = v
	}
}

// WithTimestampFormat - optional parameter to specify the timestamp format in the file - may be either a joda or java time format.
//
func (f MlFindFileStructure) WithTimestampFormat(v string) func(*MlFindFileStructureRequest) {
	return func(r *MlFindFileStructureRequest) {
		r.TimestampFormat = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f MlFindFileStructure) WithPretty() func(*MlFindFileStructureRequest) {
	return func(r *MlFindFileStructureRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f MlFindFileStructure) WithHuman() func(*MlFindFileStructureRequest) {
	return func(r *MlFindFileStructureRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f MlFindFileStructure) WithErrorTrace() func(*MlFindFileStructureRequest) {
	return func(r *MlFindFileStructureRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f MlFindFileStructure) WithFilterPath(v ...string) func(*MlFindFileStructureRequest) {
	return func(r *MlFindFileStructureRequest) {
		r.FilterPath = v
	}
}
