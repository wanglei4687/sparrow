.PHONY: go-check
go-check: fmt $(GOLANGCILINTER_BINARY)
	@$(GOLANGCILINTER_BINARY) run


.PHONY: fmt
fmt:
	@go fmt ./...

.PHONY: vet
vet:
	@go vet ./...


.PHONY: test
test: fmt vet
	@go test ./... -coverprofile cover.out
