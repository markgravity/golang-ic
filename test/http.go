package test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"

	"github.com/markgravity/golang-ic/helpers"
	"github.com/markgravity/golang-ic/helpers/log"
	
	"github.com/gin-gonic/gin"
	"github.com/onsi/ginkgo"
)

func CreateGinTestContext() (*gin.Context, *httptest.ResponseRecorder) {
	resp := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(resp)

	return c, resp
}

func MakeRequest(method string, url string, headers map[string]string, params map[string]interface{}) (*gin.Context, *httptest.ResponseRecorder) {
	if headers == nil {
		headers = map[string]string{}
	}

	headers["Content-Type"] = "application/json"

	request := buildRequest(method, url, headers, params)

	ctx, responseRecorder := CreateGinTestContext()
	ctx.Request = request

	return ctx, responseRecorder
}

func MakePostFormRequest(url string, headers map[string]string, params map[string]interface{}) (*gin.Context, *httptest.ResponseRecorder) {
	if headers == nil {
		headers = map[string]string{}
	}

	headers["Content-Type"] = "application/x-www-form-urlencoded"

	request := buildRequest(http.MethodPost, url, headers, params)

	ctx, responseRecorder := CreateGinTestContext()
	ctx.Request = request

	return ctx, responseRecorder
}

func buildRequest(method string, url string, headers map[string]string, params map[string]interface{}) *http.Request {
	request := httpRequest(method, url, headers, params)
	for k, v := range headers {
		request.Header.Set(k, v)
	}

	return request
}

func httpRequest(method string, url string, headers map[string]string, params map[string]interface{}) *http.Request {
	request, err := http.NewRequest(method, url, bodyReader(headers, params))
	if method == http.MethodGet {
		request.URL.RawQuery = helpers.GenerateURLParams(params).Encode()
	}

	if err != nil {
		ginkgo.Fail("Fail to create request: " + err.Error())
	}

	return request
}

func bodyReader(headers map[string]string, bodyData map[string]interface{}) io.Reader {
	contentType := headers["Content-Type"]

	switch contentType {
	case "application/x-www-form-urlencoded":
		data := url.Values{}
		for key, value := range bodyData {
			data.Add(key, value.(string))
		}
		return strings.NewReader(data.Encode())
	default:
		data, err := json.Marshal(bodyData)
		if err != nil {
			log.Fatal(err)
		}

		return bytes.NewReader(data)
	}
}
