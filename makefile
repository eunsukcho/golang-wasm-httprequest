wasm :
	GOOS=js GOARCH=wasm go build -o assets/json.wasm wasm/main.go