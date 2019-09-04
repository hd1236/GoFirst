package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func errPanic(writer http.ResponseWriter,
	request *http.Request) error {
	panic(123)
}

type testingUserError string

func (e testingUserError) Error() string {
	return e.Message()
}
func (e testingUserError) Message() string {
	return string(e)
}

func errUserError(writer http.ResponseWriter,
	request *http.Request) error {
	return testingUserError("user error")
}

func errNotFound(writer http.ResponseWriter,
	request *http.Request) error {
	return os.ErrNotExist
}
func errNoPermission(writer http.ResponseWriter,
	request *http.Request) error {
	return os.ErrPermission
}
func errUnknow(writer http.ResponseWriter,
	request *http.Request) error {
	return errors.New("unknow error")
}

func noError(writer http.ResponseWriter,
	request *http.Request) error {
	fmt.Fprintf(writer, "no error")
	return nil
}

//case一样
var tests = []struct {
	h       appHandler
	code    int
	message string
}{
	{errPanic, 500, "Internal Server Error"},
	{errUserError, 400, "user error"},
	{errNotFound, 404, http.StatusText(404)},
	{errNoPermission, 403, http.StatusText(403)},
	{errUnknow, 500, http.StatusText(500)},
	{noError, 200, "no error"},
}

//测试errWrapper方法
func TestErrWrapper(t *testing.T) {

	for _, tt := range tests {
		f := errWrapper(tt.h)
		resp := httptest.NewRecorder()
		req := httptest.NewRequest(
			http.MethodGet,
			"http://www.baidu.com", nil)

		f(resp, req)

		verifyResponse(resp.Result(), tt.code, tt.message, t)
	}
}

//实际运行server测试
func TestErrWrapperInServer(t *testing.T) {
	for _, tt := range tests {
		f := errWrapper(tt.h)
		server := httptest.NewServer(http.HandlerFunc(f))

		resp, _ := http.Get(server.URL)

		verifyResponse(resp, tt.code, tt.message, t)

	}
}

//提取公共代码处理response
func verifyResponse(resp *http.Response, code int, message string, t *testing.T) {
	b, _ := ioutil.ReadAll(resp.Body)
	body := strings.Trim(string(b), "\n")
	if resp.StatusCode != code ||
		body != message {
		t.Errorf("expect (%d,%s);"+
			"got (%d,%s)", code, message, resp.StatusCode, resp.Body)
	}
}
