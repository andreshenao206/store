deploy-staging: tidy format test vet lint deploySta
deploy-dev: tidy format test vet lint deployDev

tidy:
	go mod tidy
format:
	gofumpt -l -w .
build:
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/lambda lambda/*.go
test:
	go test ./... -cover
vet:
	go vet ./...
lint:
	golangci-lint run
clean:
	rm -rf ./bin

deployDev: clean build
	export AWS_PROFILE=bcode && sls deploy --stage development
deploySta: clean build
	export AWS_PROFILE=bcode && sls deploy --stage staging
remove-prod:
	export AWS_PROFILE=bcode && sls remove --stage development
