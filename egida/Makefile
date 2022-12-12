PROTO_FILES_PATH=proto
PROTO_OUT=proto

gen-proto:
	protoc -I $(PROTO_FILES_PATH) --go_out=plugins=grpc:$(PROTO_OUT) $(PROTO_FILES_PATH)/*.proto

clean-proto:
	rm -f $(PROTO_OUT)/*.pb.go

go-path:
	export GOROOT=/usr/local/go
	export GOPATH=$HOME/go
	export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
	export CGO_ENABLED=0

build: go-path
	go build -o build/egida cmd/main.go

docs:
	mkdocs gh-deploy

dsl:
	java -jar tools/antlr-4.8.jar -o pkg/parser -listener -visitor -package parser -Dlanguage=Go Aspida.g4