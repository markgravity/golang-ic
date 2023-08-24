package test

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/onsi/ginkgo"
	"gopkg.in/dnaeon/go-vcr.v3/cassette"
	"gopkg.in/dnaeon/go-vcr.v3/recorder"
)

func CassettePath(cassetteName string) string {
	return fmt.Sprintf("%s/test/fixtures/vcr/%s", RootDir(), cassetteName)
}

func GetRecorderClient(cassetteName string) (*http.Client, *recorder.Recorder) {
	rec, err := recorder.New(CassettePath(cassetteName))
	if err != nil {
		ginkgo.Fail(err.Error())
	}

	if rec.Mode() != recorder.ModeRecordOnce {
		ginkgo.Fail("Recorder should be in ModeRecordOnce")
	}

	rec.SetMatcher(customMatcher)

	return rec.GetDefaultClient(), rec
}

func customMatcher(request *http.Request, recordedReq cassette.Request) bool {
	if request.Body == nil || request.Body == http.NoBody {
		return cassette.DefaultMatcher(request, recordedReq)
	}

	reqBody, err := io.ReadAll(request.Body)
	if err != nil {
		ginkgo.Fail("Fail to read request body")
	}
	request.Body.Close()
	request.Body = io.NopCloser(bytes.NewBuffer(reqBody))

	isMatchRequestBody := string(reqBody) == recordedReq.Body

	isMethodMatch := request.Method == recordedReq.Method
	isURLMatch := request.URL.String() == recordedReq.URL

	return isMethodMatch && isURLMatch && isMatchRequestBody
}
