VERSION=$(shell git rev-parse --short HEAD)

build-dev:
	docker build --build-arg APP_VERSION=$(VERSION) .