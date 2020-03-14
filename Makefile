.PHONY: proto test datastore-start datastore-stop install install-drone run build ui

export GIZMO_SKIP_OBSERVE=true;
export CORS_ALLOWED_ORIGIN=http://localhost:3000;

proto:
	@rm -f ./service.proto ./service.pb ./service.pb.go;
	@openapi2proto -spec service.yaml -out ./service.proto;
	@protoc --js_out=import_style=commonjs,binary:. --go_out=plugins=grpc:. ./service.proto;
	@mv ./service_pb.js ./ui;

test: 
	@goimports -w ./**/**/*.go;
	@gofmt -s -w ./**/**/*.go;
	@go vet ./...;
	@go test -race -cover -vet=off ./...;

run:
	@go run -race ./cmd/server;

build:
	@go build ./cmd/server;

datastore-start: datastore-stop
	@gcloud beta emulators datastore start --no-store-on-disk --quiet > /dev/null 2>&1 &
	@gcloud beta emulators datastore env-init --quiet > /dev/null 2>&1;

datastore-stop:
	@kill -9 `ps ax | grep 'CloudDatastore.jar' | grep -v grep | awk '{print $1}'` > /dev/null 2>&1;

install:
	@brew cask install google-cloud-sdk;
	@gcloud init;
	@gcloud config set project financial-calculator-dev;
	@gcloud components install cloud-datastore-emulator;
	@go get -u github.com/NYTimes/openapi2proto/cmd/openapi2proto;

install-drone:
	@go get golang.org/x/tools/cmd/goimports;

test-drone:
	@goimports -w ./**/**/*.go;
	@gofmt -s -w ./**/**/*.go;
	@go vet ./...;
	@go test -race -cover -vet=off -coverprofile=coverage.txt -covermode=atomic ./...;

ui:
	@cd ./ui; \
	npm start;

build-ui:
	@cd ./ui; \
	npm ci; \
	npm run build;	
