// Code generated by depstubber. DO NOT EDIT
// This is a simple stub for github.com/astaxie/beego/context, strictly for use in testing.

// See the LICENSE file for information about the licensing of the original library.
// Source: github.com/astaxie/beego/context (exports: BeegoInput,BeegoOutput,Context; functions: WriteBody)

// Package context is a stub of github.com/astaxie/beego/context, generated by depstubber.
package context

import (
	"bufio"
	"io"
	"net"
	"net/http"
	"reflect"
	"time"
)

type BeegoInput struct {
	Context       *Context
	CruSession    interface{}
	RequestBody   []byte
	RunMethod     string
	RunController reflect.Type
}

func (_ *BeegoInput) AcceptsHTML() bool {
	return false
}

func (_ *BeegoInput) AcceptsJSON() bool {
	return false
}

func (_ *BeegoInput) AcceptsXML() bool {
	return false
}

func (_ *BeegoInput) AcceptsYAML() bool {
	return false
}

func (_ *BeegoInput) Bind(_ interface{}, _ string) error {
	return nil
}

func (_ *BeegoInput) Cookie(_ string) string {
	return ""
}

func (_ *BeegoInput) CopyBody(_ int64) []byte {
	return nil
}

func (_ *BeegoInput) Data() map[interface{}]interface{} {
	return nil
}

func (_ *BeegoInput) Domain() string {
	return ""
}

func (_ *BeegoInput) GetData(_ interface{}) interface{} {
	return nil
}

func (_ *BeegoInput) Header(_ string) string {
	return ""
}

func (_ *BeegoInput) Host() string {
	return ""
}

func (_ *BeegoInput) IP() string {
	return ""
}

func (_ *BeegoInput) Is(_ string) bool {
	return false
}

func (_ *BeegoInput) IsAjax() bool {
	return false
}

func (_ *BeegoInput) IsDelete() bool {
	return false
}

func (_ *BeegoInput) IsGet() bool {
	return false
}

func (_ *BeegoInput) IsHead() bool {
	return false
}

func (_ *BeegoInput) IsOptions() bool {
	return false
}

func (_ *BeegoInput) IsPatch() bool {
	return false
}

func (_ *BeegoInput) IsPost() bool {
	return false
}

func (_ *BeegoInput) IsPut() bool {
	return false
}

func (_ *BeegoInput) IsSecure() bool {
	return false
}

func (_ *BeegoInput) IsUpload() bool {
	return false
}

func (_ *BeegoInput) IsWebsocket() bool {
	return false
}

func (_ *BeegoInput) Method() string {
	return ""
}

func (_ *BeegoInput) Param(_ string) string {
	return ""
}

func (_ *BeegoInput) Params() map[string]string {
	return nil
}

func (_ *BeegoInput) ParamsLen() int {
	return 0
}

func (_ *BeegoInput) ParseFormOrMulitForm(_ int64) error {
	return nil
}

func (_ *BeegoInput) Port() int {
	return 0
}

func (_ *BeegoInput) Protocol() string {
	return ""
}

func (_ *BeegoInput) Proxy() []string {
	return nil
}

func (_ *BeegoInput) Query(_ string) string {
	return ""
}

func (_ *BeegoInput) Refer() string {
	return ""
}

func (_ *BeegoInput) Referer() string {
	return ""
}

func (_ *BeegoInput) Reset(_ *Context) {}

func (_ *BeegoInput) ResetParams() {}

func (_ *BeegoInput) Scheme() string {
	return ""
}

func (_ *BeegoInput) Session(_ interface{}) interface{} {
	return nil
}

func (_ *BeegoInput) SetData(_ interface{}, _ interface{}) {}

func (_ *BeegoInput) SetParam(_ string, _ string) {}

func (_ *BeegoInput) Site() string {
	return ""
}

func (_ *BeegoInput) SubDomains() string {
	return ""
}

func (_ *BeegoInput) URI() string {
	return ""
}

func (_ *BeegoInput) URL() string {
	return ""
}

func (_ *BeegoInput) UserAgent() string {
	return ""
}

type BeegoOutput struct {
	Context    *Context
	Status     int
	EnableGzip bool
}

func (_ *BeegoOutput) Body(_ []byte) error {
	return nil
}

func (_ *BeegoOutput) ContentType(_ string) {}

func (_ *BeegoOutput) Cookie(_ string, _ string, _ ...interface{}) {}

func (_ *BeegoOutput) Download(_ string, _ ...string) {}

func (_ *BeegoOutput) Header(_ string, _ string) {}

func (_ *BeegoOutput) IsCachable() bool {
	return false
}

func (_ *BeegoOutput) IsClientError() bool {
	return false
}

func (_ *BeegoOutput) IsEmpty() bool {
	return false
}

func (_ *BeegoOutput) IsForbidden() bool {
	return false
}

func (_ *BeegoOutput) IsNotFound() bool {
	return false
}

func (_ *BeegoOutput) IsOk() bool {
	return false
}

func (_ *BeegoOutput) IsRedirect() bool {
	return false
}

func (_ *BeegoOutput) IsServerError() bool {
	return false
}

func (_ *BeegoOutput) IsSuccessful() bool {
	return false
}

func (_ *BeegoOutput) JSON(_ interface{}, _ bool, _ bool) error {
	return nil
}

func (_ *BeegoOutput) JSONP(_ interface{}, _ bool) error {
	return nil
}

func (_ *BeegoOutput) Reset(_ *Context) {}

func (_ *BeegoOutput) ServeFormatted(_ interface{}, _ bool, _ ...bool) {}

func (_ *BeegoOutput) Session(_ interface{}, _ interface{}) {}

func (_ *BeegoOutput) SetStatus(_ int) {}

func (_ *BeegoOutput) XML(_ interface{}, _ bool) error {
	return nil
}

func (_ *BeegoOutput) YAML(_ interface{}) error {
	return nil
}

type Context struct {
	Input          *BeegoInput
	Output         *BeegoOutput
	Request        *http.Request
	ResponseWriter *Response
}

func (_ *Context) Abort(_ int, _ string) {}

func (_ *Context) CheckXSRFCookie() bool {
	return false
}

func (_ *Context) GetCookie(_ string) string {
	return ""
}

func (_ *Context) GetSecureCookie(_ string, _ string) (string, bool) {
	return "", false
}

func (_ *Context) Redirect(_ int, _ string) {}

func (_ *Context) RenderMethodResult(_ interface{}) {}

func (_ *Context) Reset(_ http.ResponseWriter, _ *http.Request) {}

func (_ *Context) SetCookie(_ string, _ string, _ ...interface{}) {}

func (_ *Context) SetSecureCookie(_ string, _ string, _ string, _ ...interface{}) {}

func (_ *Context) WriteString(_ string) {}

func (_ *Context) XSRFToken(_ string, _ int64) string {
	return ""
}

type Response struct {
	ResponseWriter http.ResponseWriter
	Started        bool
	Status         int
	Elapsed        time.Duration
}

func (_ Response) Header() http.Header {
	return nil
}

func (_ *Response) CloseNotify() <-chan bool {
	return nil
}

func (_ *Response) Flush() {}

func (_ *Response) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	return nil, nil, nil
}

func (_ *Response) Pusher() http.Pusher {
	return nil
}

func (_ *Response) Write(_ []byte) (int, error) {
	return 0, nil
}

func (_ *Response) WriteHeader(_ int) {}

func WriteBody(_ string, _ io.Writer, _ []byte) (bool, string, error) {
	return false, "", nil
}