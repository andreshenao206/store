package routers

import "net/http"

// LambdaResponse response for the AWS Lambda Request functionality.
type LambdaResponse struct {
	Error  string      `json:"error,omitempty"`
	Result interface{} `json:"result,omitempty"`
}

// LuvEvent request for the AWS Lambda Request functionality.
type LuvEvent struct {
	RouterName  string            		`json:"routerName"`
	Authorizer  map[string]interface{}  `json:"authorizer"`
	Payload     map[string]string 		`json:"payload"`
}

// LambdaClientError to build and return a Response Struct for client errors.
func LambdaClientError(status int) (LambdaResponse, error) {
	return LambdaResponse{
		Error: http.StatusText(status),
	}, nil
}

// LambdaServerError to build and return a Response Struct for application errors.
func LambdaServerError(err error) (LambdaResponse, error) {
	return LambdaResponse{
		Error: err.Error(),
	}, nil
}

// LambdaSuccessResponse to build and return a Response Struct for successful responses.
func LambdaSuccessResponse(result interface{}) LambdaResponse {
	return LambdaResponse{
		Result: result,
	}
}
