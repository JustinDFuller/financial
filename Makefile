.PHONY: test datastore-start install

proto:
	@openapi2proto -spec types/service.yaml -out types/service.proto;
	@protoc --go_out=plugins=grpc:. types/service.proto;
	@protoc --include_imports --include_source_info types/service.proto --descriptor_set_out types/service.pb;

test: datastore-start
	@goimports -w ./**/*.go;
	@gofmt -s -w ./**/*.go;
	@go test -race ./...;
	@$(MAKE) datastore-stop;

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

