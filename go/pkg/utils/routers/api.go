package routers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

// APIResponse is of type APIGatewayProxyResponse since we're leveraging the AWS Lambda Proxy Request functionality.
type APIResponse events.APIGatewayProxyResponse
type ErrorInfo struct {
	Message string `json:"message"`
}
type ErrorResponse struct {
	Error ErrorInfo `json:"error"`
}

type ResultInfo struct {
	Message string `json:"message"`
	Id string `json:"id,omitempty"`
}
type ResultResponse struct {
	Result ResultInfo `json:"result"`
}

var ErrNoRouterCaseImpl = errors.New("no router case implemented for this endpoint")

// APIServerError to build and return a Response Struct for application errors.
func APIServerError(e error) (response APIResponse, err error) {
	eBytes, err := json.Marshal(ErrorResponse{Error: ErrorInfo{Message: e.Error()}})
	if err != nil {
		return
	}
	response = APIResponse{
		StatusCode: http.StatusInternalServerError,
		Body:       string(eBytes),
		Headers:    map[string]string{"Content-Type": "application/json"},
	}
	return
}

// APISuccessResponse to build and return a Response Struct for successful responses.
func APISuccessResponse(body string) APIResponse {
	return APIResponse{
		StatusCode: 200,
		Body:       body,
		Headers:    map[string]string{"Content-Type": "application/json"},
	}
}

// HandleResponse to handle the response/
func HandleResponse(message string, err error) (APIResponse, error) {
	if err != nil {
		log.Printf(":::Error::: %s", err.Error())
		return APIServerError(err)
	}
	return APISuccessResponse(message), nil
}