package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/store/go/userprofile/lambda/config"
	"github.com/store/go/userprofile/lambda/routers"
	"github.com/store/go/userprofile/lambda/usecases"

	"github.com/aws/aws-lambda-go/lambda"
)

func Alive() string {
	return "{\"status\":\"ok\"}"
}

const (
	HTTP_METHOD_GET  string = "GET"
	HTTP_METHOD_POST string = "POST"
)

var appConfig = config.LoadConfig()

func router(request events.APIGatewayProxyRequest) (routers.APIResponse, error) {
	switch request.HTTPMethod {
	case HTTP_METHOD_GET:
		message := Alive()
		return routers.APISuccessResponse(message), nil
		//case path.Method crud
	case HTTP_METHOD_POST:
		switch path := request.Path; {
		case path == appConfig.CreateUser:
			return routers.HandleResponse(usecases.CreateUser(request))

		default:
			return routers.APIServerError(routers.ErrNoRouterCaseImpl)
		}

	default:
		err := errors.New(http.StatusText(http.StatusMethodNotAllowed))
		log.Print(err)
		return routers.APIServerError(err)
	}
}

func main() {
	lambda.Start(router)
}
