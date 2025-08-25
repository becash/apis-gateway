
.PHONY: update_modules
update_modules:
	@go get -u all
	@go mod tidy
	@go mod vendor


.PHONY: update_vendor
update_vendor:
	@go get -u github.com/becash/apis
	@go mod tidy
	@go mod vendor


.PHONY: generate
generate:
	@#gqlgen generate
	@go generate ./...