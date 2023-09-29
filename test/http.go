package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/markgravity/golang-ic/lib/models"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"net/textproto"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/markgravity/golang-ic/helpers/log"
	"github.com/markgravity/golang-ic/test/helpers"

	"github.com/gin-gonic/gin"
	"github.com/google/jsonapi"
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

func MakeAuthenticatedRequest(method string, url string, headers map[string]string, params map[string]interface{}, user *models.User) (*gin.Context, *httptest.ResponseRecorder) {
	if headers == nil {
		headers = map[string]string{}
	}
	accessToken := "Bearer " + helpers.GenerateToken(user.Base.ID.String())
	headers["Authorization"] = accessToken

	return MakeRequest(method, url, headers, params)
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

func MakeMultipartRequestRequest(url string, filePath string, contentType string, headers http.Header, user *models.User) (*gin.Context, *httptest.ResponseRecorder) {
	if user != nil {
		if headers == nil {
			headers = http.Header{}
		}

		headers.Set(
			"Authorization",
			"Bearer "+helpers.GenerateToken(user.Base.ID.String()),
		)
	}

	newHeaders, payload := CreateMultipartRequestInfo(filePath, contentType, headers)
	request, _ := http.NewRequest("POST", url, payload)
	request.Header = newHeaders

	ctx, responseRecorder := CreateGinTestContext()
	ctx.Request = request

	return ctx, responseRecorder
}

func GetMultipartAttributesFromFile(filePath string, contentType string) (multipart.File, *multipart.FileHeader, error) {
	headers, payload := CreateMultipartRequestInfo(filePath, contentType, nil)
	req, err := http.NewRequest("POST", "", payload)
	if err != nil {
		return nil, nil, err
	}

	req.Header = headers
	file, fileHeader, err := req.FormFile("file")

	return file, fileHeader, err
}

func CreateMultipartRequestInfo(filePath string, contentType string, headers http.Header) (http.Header, *bytes.Buffer) {
	realPath := fmt.Sprintf("%s/test/fixtures/files/%s", RootDir(), filePath)
	file, err := os.Open(realPath)
	if err != nil {
		log.Error("Failed to open file: ", err)
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := createFormFile(writer, "file", filepath.Base(filePath), contentType)
	if err != nil {
		log.Error("Failed to create part from file: ", err)
	}

	_, err = io.Copy(part, file)
	if err != nil {
		log.Error("Failed to copy file: ", err)
	}
	writer.Close()

	if headers == nil {
		headers = http.Header{}
	}
	headers.Set("Content-Type", writer.FormDataContentType())

	return headers, body
}

func createFormFile(w *multipart.Writer, fieldname string, filePath string, contentType string) (io.Writer, error) {
	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s"; filename="%s"`, fieldname, filePath))
	h.Set("Content-Type", contentType)
	return w.CreatePart(h)
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

func UnmarshalErrorResponseBody(responseBody *bytes.Buffer) jsonapi.ErrorsPayload {
	response := jsonapi.ErrorsPayload{}
	err := json.Unmarshal(responseBody.Bytes(), &response)
	if err != nil {
		ginkgo.Fail("Fail to unmarshal response body" + err.Error())
	}

	return response
}
