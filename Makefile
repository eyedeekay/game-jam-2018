
GOPATH=$(shell pwd)/.go

export ANDROID_HOME=$(shell pwd)/.android

GOPHERJS=$(GOPATH)/bin/gopherjs

GOMOBILE=$(GOPATH)/bin/gomobile

APPLICATION=gamewrapper
#ALPINE=-alpine
#NETGO=--tags netgo

ANDROID_HOME=/opt/android-sdk-linux
SDK_TOOLS_VERSION=4333796
#ANDROID_EXTRAS="add-ons"
CONTAINER_PATH=$(PATH):$(ANDROID_HOME)/tools:$(ANDROID_HOME)/tools/bin:$(ANDROID_HOME)/platform-tools

build: bin fmt
	go build $(NETGO) -o bin/demo

run: build
	./bin/demo -width 1024 -height 768

android-build:
	gomobile build -tags android -target android -o bin/demo.apk

ios-build:
	gomobile build -target ios -o bin/demo.app

clean:
	rm *.i2pkeys bin/demo www/demo.js www/demo.js.map -f

alldeps: deps engo android gomobile-init engo-android

deps:
	go get -u github.com/mattn/anko
	go get -u github.com/eyedeekay/sam-forwarder/config

engo: deps
	go get -u engo.io/engo/common
	go get -u engo.io/engo
	go get -u engo.io/ecs

engo-static:
	go get -u -tags netgo engo.io/engo

android:
	go get -u golang.org/x/mobile/cmd/gomobile
	go get -u golang.org/x/mobile/cmd/gobind

engo-android: android
	go get -u -tags android engo.io/engo

gomobile-init: .android
	$(GOMOBILE) init

engo-mobile: android engo-android gomobile-init

fmt:
	find . -path ./.go -prune -o -name "*.go" -exec gofmt -w {} \;

www:
	mkdir -p www

bin:
	mkdir -p bin

.android/ndk-bundle:
	mkdir -p .android/ndk-bundle

js: www/demo.js

www/demo.js:
	$(GOPHERJS) build $(NETGO) -o www/demo.js

rejs: www
	$(GOPHERJS) build $(NETGO) -o www/demo.js

html: www
	@echo "<!doctype html>"
	@echo "<html lang=en>"
	@echo "  <head>"
	@echo "    <meta charset=utf-8>"
	@echo "    <title>$(APPLICATION)</title>"
	@echo "    <script src=\"demo.js\"> </script>"
	@echo "  </head>"
	@echo "  <body>"
	@echo "  </body>"
	@echo "</html>"

index: www rejs
	make -s html | tee index.html

gopherjs:
	go get -u github.com/gopherjs/gopherwasm/js
	go get -u github.com/gopherjs/gopherjs

serve:
	cd www && $(GOPHERJS) serve ../

i2pserve:
	eephttpd; rm -f eephttpd.i2pkeys

install:
	go install ./config ./entity ./game ./graphic ./net ./system

docker:
	docker build \
		--build-arg ANDROID_HOME=$(ANDROID_HOME) \
		--build-arg SDK_TOOLS_VERSION=$(SDK_TOOLS_VERSION) \
		--build-arg API_LEVELS=$(API_LEVELS) \
		--build-arg BUILD_TOOLS_VERSIONS=$(BUILD_TOOLS_VERSIONS) \
		--build-arg ANDROID_EXTRAS=$(ANDROID_EXTRAS) \
		--build-arg PATH=$(CONTAINER_PATH) \
		-t eyedeekay/$(APPLICATION) .

docker-run:
	docker run -ti --rm \
		-e ANDROID_HOME=$(ANDROID_HOME) \
		-e SDK_TOOLS_VERSION=$(SDK_TOOLS_VERSION) \
		-e API_LEVELS=$(API_LEVELS) \
		-e BUILD_TOOLS_VERSIONS=$(BUILD_TOOLS_VERSIONS) \
		-e ANDROID_EXTRAS=$(ANDROID_EXTRAS) \
		-e PATH=$(PATH) \
		-e DISPLAY=$(DISPLAY) \
		--device /dev/snd \
		-v /tmp/.X11-unix:/tmp/.X11-unix \
		-v $(HOME)/.Xauthority:/home/gamewrapper/.Xauthority \
		eyedeekay/$(APPLICATION)$(ALPINE)

docker-alpine:
	docker build \
		--build-arg ANDROID_HOME=$(ANDROID_HOME) \
		--build-arg SDK_TOOLS_VERSION=$(SDK_TOOLS_VERSION) \
		--build-arg API_LEVELS=$(API_LEVELS) \
		--build-arg BUILD_TOOLS_VERSIONS=$(BUILD_TOOLS_VERSIONS) \
		--build-arg ANDROID_EXTRAS=$(ANDROID_EXTRAS) \
		--build-arg PATH=$(CONTAINER_PATH) \
		-f Dockerfile.alpine \
		-t eyedeekay/$(APPLICATION)-alpine .
