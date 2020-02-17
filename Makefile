.PHONY: proto test datastore-start datastore-stop install run

export GIZMO_SKIP_OBSERVE=true;

proto:
	@rm -f service/service.proto service/service.pb service/service.pb.go;
	@openapi2proto -spec service/service.yaml -out service/service.proto;
	@protoc --go_out=plugins=grpc:. service/service.proto;
	@protoc --include_imports --include_source_info service/service.proto --descriptor_set_out service/service.pb;

test: 
	@goimports -w ./**/*.go;
	@gofmt -s -w ./**/*.go;
	@go test -race -cover ./...;

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

