Dapr-WASM enables you to run WASM code in Dapr.

> This is a very early preview of the experimental code to run wasm code using dapr capabilities, and you should expect breaking changes before the first stable release.

### Project description
This project is based on dapr and adds wasm support to run wasm code in dapr. wasm code can use dapr capabilities, such as service call, pubsub, etc.

### Architecture
! [image-20230305161929424](https://image-1255620078.cos.ap-nanjing.myqcloud.com/image-20230305161929424.png)

### Advantages
1. no need to write any glue code, just wasm function code
2. wasm is lightweight， security, isolation and has other features
3. allow wasm code to use dapr capabilities such as service call, pubsub, etc. So wasm applications can focus entirely on business logic

### eamples
In the project `examples` directory, there is a simple example that can receive network requests and use state store to store and fetch data.
##### 1. Compile the wasm code
```shell
cd examples/service-call
make
```
##### 2. Run dapr
Install and initialize dapr

Run the `make run-example` command in the home directory
##### 3. Send the request
```shell
# Set the value
curl http://localhost:3500/v1.0/invoke/go-wasm-app/method/service-call/set -d 'value1'
# Query, this should return value1
curl http://localhost:3500/v1.0/invoke/go-wasm-app/method/service-call/get
```

### Next step
Refine features and details
Add wasm instance pool, which can reuse wasm instance as well as parallelize requests
Now it is wasm->wasm-dapr(wasm host + dapr sdk) -> dapr,later will explore directly on dapr long, wasm->dapr.