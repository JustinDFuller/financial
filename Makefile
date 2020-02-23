.PHONY: proto test datastore-start datastore-stop install run

export GIZMO_SKIP_OBSERVE=true;

proto:
	@rm -f ./service.proto ./service.pb ./service.pb.go;
	@openapi2proto -spec service.yaml -out ./service.proto;
	@protoc --go_out=plugins=grpc:. ./service.proto;

test: 
	@goimports -w ./**/**/*.go;
	@gofmt -s -w ./**/**/*.go;
	@go vet ./...;
	@go test -race -cover -vet=off ./...;

run:
	@go run -race ./cmd/server;

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

