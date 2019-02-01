IMAGE = github.com/smgladkovskiy/structs
TEST_CONTAINER = docker run --rm -i --name nulls_test $(IMAGE):test

deps: ## Get dependencies
	@dep ensure

race: ## Run data race detector
	@go test -race ./...

test: ## Run application unit tests with coverage and generate global code coverage report
	@go test ./... -parallel 4 -failfast -cover -coverprofile=.test_artifacts/coverage.out -bench=. -benchmem

covercli: ## Generate code coverage report
	@go tool cover -func=.test_artifacts/coverage.out

coverhtml: ## Generate global code coverage report in HTML
	@go tool cover -html=.test_artifacts/coverage.out

coverage: test coverhtml

coverage_cli: test covercli

image_test: ## Create test image
	@docker build -t $(IMAGE):test .

test_in_docker: image_test ## Testing code with unit tests in docker container
	@$(TEST_CONTAINER) make coverage_cli

all: deps image_test
