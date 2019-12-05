.PHONY: build clean fmt scp run

scp: build
	@echo "scp"
	@scp ./build/pi pi@192.168.50.79:~

run: scp
	@echo "run"
	@ssh pi@192.168.50.79 "./pi"

build: clean
	@echo "build"
	GOARM=7 GOARCH=arm GOOS=linux go build -o build/pi  cmd/pi/main.go

clean:
	@echo "clean"
	@rm -rf build/*

fmt:
	@find . -name "*.go" -type f -exec go fmt {} \;
