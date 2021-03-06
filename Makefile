
.PHONY: install-pre-commit
install-pre-commit:
	python3 -m pip install --user --upgrade pip ; python3 -m pip install --user pre-commit gitlint
	pre-commit install -f ; pre-commit install --hook-type commit-msg

.PHONY: windows
windows:
	bash curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s -- -b %GOPATH%/bin

.PHONY: linux
linux:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

.PHONY: linters
linters:
	golangci-lint run ./codewar/golang/...
	golangci-lint run ./golang/...

.PHONY: tests
tests:
	go test -short -timeout 30s -cover ./codewar/golang/...
	go test -short -timeout 30s -cover ./golang/...