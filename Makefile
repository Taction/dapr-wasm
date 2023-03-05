
build:
	go build -o dwasm main.go

run-example:
	dapr run --app-id go-wasm-app --app-port 6003 --app-protocol http --dapr-http-port 3500 -- go run main.go --port 6003 --config config.example.yaml