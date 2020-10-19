build:
	rm -r docs/* || true
	GOOS=js GOARCH=wasm go build -o docs/starling.wasm
	cp $$(go env GOROOT)/misc/wasm/wasm_exec.js docs/
	cp static/* docs/

serve: build
	go run cmd/serve.go
