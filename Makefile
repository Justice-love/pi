.PHONY: scp run build clean fmt vet test gen

PKGS := $(shell go list ./... )

scp: build
	@echo "scp"
	@scp ./build/pi pi@192.168.50.79:~

run: scp
	@echo "run"
	@ssh pi@192.168.50.79 "./pi"

build: clean
	@echo "build"
	GOARM=7 GOARCH=arm GOOS=linux go build -o build/pi  cmd/car/main.go

clean:
	@echo "clean"
	@rm -rf build/*

fmt:
	@find . -name "*.go" -type f -exec go fmt {} \;

vet:
	go vet ${PKGS}

test:
	@go test -p 1 -v -cover ${PKGS}

gen:
	@go generate internal/car/api.go
