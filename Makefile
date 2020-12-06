
unittests:
	go test ./... -count=1

installation-cleanup:
	@rm -r /tmp/gen

install:
	@GO111MODULE=off go get k8s.io/gengo/...
	@GO111MODULE=off go get github.com/Prepodavan/gombok
	@export PWD=$$(pwd) && \
		cd / && \
		GO111MODULE=on go get k8s.io/gengo/... && \
		cd $$PWD
	@export PWD=$$(pwd) && \
		git clone https://github.com/Prepodavan/gen /tmp/gen && \
		cd /tmp/gen && \
	 	go build -o $$(go env GOPATH)/bin && \
	 	cd $$PWD

generate:
	@go generate ./...
	@deepcopy-gen --logtostderr --v=1 -i ./requests,./models,./responses -o .
	@gen ./models ./responses ./requests
