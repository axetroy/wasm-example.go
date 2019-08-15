### WebAssembly example writing in Golang

#### Example: hello
```bash
# build wasm
GOARCH=wasm GOOS=js go build -o ./demo/test.wasm ./hello/main.go
# run in server
go run demo.go
# Open http://localhost:8080 in your browser
```

#### Example: fibonacci
```bash
# build wasm
GOARCH=wasm GOOS=js go build -o ./demo/test.wasm ./fibonacci/main.go
# run in server
go run demo.go
# Open http://localhost:8080 in your browser
```