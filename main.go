package main

import (
	"fmt"
	"github.com/ysyesilyurt/go-restclient/restclient"
	"net/http"
	"time"
	"zap-fieldbuilder-template/logger"
)

/* Driver code for Generic uber-go/zap FieldBuilder Template - ysyesilyurt 2021 */

func main() {
	/* By the way, go-restclient is a handy Go Package to Call Internal HTTP APIs with Builder Pattern (as this one :))
	check https://github.com/ysyesilyurt/go-restclient for further information (A tiny visit is highly recommended :)) */

	type inlineDummyRequestBody struct {
		DummyPayload string `json:"dummyPayload"`
	}

	type inlineDummyResponseBody struct {
		DummyPayload string `json:"dummyPayload"`
	}

	reqBody := inlineDummyRequestBody{DummyPayload: "dummy request"}
	respBody := inlineDummyResponseBody{DummyPayload: "dummy response"}
	rawUrl := fmt.Sprintf("https://ysyesilyurt.com/posts/%d/comments/%d?from=%s&since=%s", 123, 456, "johndoe", time.Now().Add(-2*time.Hour).Format(time.RFC3339))
	req, _ := restclient.RequestBuilder().
		RawUrl(rawUrl).
		Header(&http.Header{"Content-Type": []string{"application/json"}, "Cookie": []string{"test-1234"}}).
		BodyJson(reqBody).
		ResponseReference(respBody).
		LoggingEnabled(false).
		Build()

	log := logger.GetContextLogger(nil)
	start := time.Now()

	requestLogFields := logger.FieldBuilder().
		Issuer("ysyesilyurt").
		AccountId(123).
		Url(rawUrl).
		Method(http.MethodPost).
		Timestamp(start.Format(time.RFC3339))

	log.Debug("Sending Awesome Request...", requestLogFields.Build()...)

	reqErr := req.Post()

	end := time.Now()
	duration := end.Sub(start)

	responseLogFields := requestLogFields.
		Status(reqErr.GetStatusCode()).
		DurationMs(duration.Milliseconds()).
		Timestamp(end.Format(time.RFC3339))

	if reqErr != nil {
		log.Info("Oops! Request failed.", responseLogFields.Build()...)
	} else {
		log.Info("Yaay! Request completed.", responseLogFields.Build()...)
	}
}
