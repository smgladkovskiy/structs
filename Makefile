IMAGE = rg.teamc.io/teamc.io/golang/nulls
TEST_CONTAINER = docker run --rm -i --name nulls_test $(IMAGE):test

deps: ## Get and update dependencies
	@go get -t

test: ## Run application unit tests with coverage and generate global code coverage report
	@go test ./... -parallel 4 -failfast -cover -coverprofile=coverage.out

covercli: ## Generate code coverage report
	@go tool cover -func=coverage.out

image_test:
	@docker build -t $(IMAGE):test .

test_in_docker: image_test ## Testing code with unit tests in docker container
	@$(TEST_CONTAINER) make coverage_cli

coverage_cli: test covercli
