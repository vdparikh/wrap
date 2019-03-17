package wrap

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/labstack/echo"
)

func formatAPIResponse(statusCode int, headers http.Header, responseData string) (events.APIGatewayProxyResponse, error) {
	responseHeaders := make(map[string]string)

	responseHeaders["Content-Type"] = "application/json"
	for key, value := range headers {
		responseHeaders[key] = ""

		if len(value) > 0 {
			responseHeaders[key] = value[0]
		}
	}

	responseHeaders["Access-Control-Allow-Origin"] = "*"
	responseHeaders["Access-Control-Allow-Headers"] = "origin,Accept,Authorization,Content-Type"

	return events.APIGatewayProxyResponse{
		Body:       responseData,
		Headers:    responseHeaders,
		StatusCode: statusCode,
	}, nil
}

// Route wraps echo server into Lambda Handler
func Route(e *echo.Echo) func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		body := strings.NewReader(request.Body)
		req := httptest.NewRequest(request.HTTPMethod, request.Path, body)
		for k, v := range request.Headers {
			req.Header.Add(k, v)
		}

		q := req.URL.Query()
		for k, v := range request.QueryStringParameters {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()

		rec := httptest.NewRecorder()
		e.ServeHTTP(rec, req)

		res := rec.Result()
		responseBody, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return formatAPIResponse(http.StatusInternalServerError, res.Header, err.Error())
		}

		return formatAPIResponse(res.StatusCode, res.Header, string(responseBody))
	}
}
