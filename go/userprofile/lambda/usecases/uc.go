package usecases

import (
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/store/go/userprofile/lambda/adapters"
	"github.com/store/go/userprofile/lambda/config"
	"github.com/store/go/userprofile/lambda/domain"
)

var appConfig = config.LoadConfig()

func CreateUser(request events.APIGatewayProxyRequest) (str string, err error) {
	var requestPost domain.User
	err = json.Unmarshal([]byte(request.Body), &requestPost)
	if err != nil {
		log.Printf("::::POST:::: request post body err  %v", err)
		return
	}
	log.Printf("::::POST:::: request post body user %v", requestPost)

	adapters.RequestNode(requestPost, config.User)

	return
}
