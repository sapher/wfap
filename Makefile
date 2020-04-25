clean:
	@rm -rf build

build: clean
	@mkdir build
	go build -i -o ./build/wfap ./cmd/wfap