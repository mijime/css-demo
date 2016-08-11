BUILD_TARGET = \
							 server/bindata.go \
							 build/Release/css-demo \
							 build/Release/assets

GOSRC = $(wildcard *.go */*.go)

build: $(BUILD_TARGET)

serv:
	make build
	build/Release/css-demo --debug --addr :3000

setup:
	npm install

build/Release/css-demo: $(GOSRC) server/bindata.go
	go fmt ./...
	go get -v ./...
	go build -o $@

server/bindata.go: build/Release/assets templates
	go get -v github.com/jteeuwen/go-bindata/...
	go-bindata -pkg server -o $@ $^

build/Release/assets:
	npm run build -- --output-path $@

clean:
	rm $(BUILD_TARGET)
