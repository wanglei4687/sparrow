.PHONY: build
build:
	@go build -o ./bin/sparrow ./cmd/sparrow/main.go

.PHONY:	docker-build
docker-build: fmt vet
	@docker build  -t $(IMG_TAG)  .
