
# Wrap
A very simple wrapper to wrap your [Labstack Echo ](https://github.com/labstack/echo) API's into Lambda/API Gateway

### Installation
```
go get -u github.com/vdparikh/wrap
go mod vendor
```

### Usage 
Check out example/example.go for usage.

```go
package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/vdparikh/wrap"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())

	e.GET("/hello", func(c echo.Context) error {
		return c.JSON(200, "HELLO")
	})

	server := wrap.Route(e)

	lambda.Start(server)
}
```

### Run sample

To run this locally as SAM (Serverless Application Model), then you need to install SAM CLI. SAM is an easier way to test out your serverless functions without having to deploy them first.
- [AWS SAM CLI](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/serverless-sam-cli-install-mac.html)

```sh
cd example
# Build Binary
GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o hello .

# Start SAM
sam local start-api
```

On another terminal execute the API and you should see the HELLO response
```sh
curl -v http://localhost:3000/hello
```