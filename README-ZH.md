Dapr-WASM enables you to run WASM code in Dapr.

> This is a very early preview of the experimental code to run wasm code using dapr capabilities, and you should expect breaking changes before the first stable release.

### 项目介绍
本项目是在dapr的基础上，增加了wasm的支持，可以在dapr中运行wasm代码，wasm代码可以使用dapr的能力，比如调用其他服务，发布消息等。

### 架构
![image-20230305161929424](https://image-1255620078.cos.ap-nanjing.myqcloud.com/image-20230305161929424.png)

### 优势
1. 无需编写任何胶水代码，只需要编写wasm功能代码即可
2. wasm非常轻量并且具有安全、隔离性等特性
3. 让wasm代码可以使用dapr的能力，比如调用其他服务，发布消息等。因此wasm程序可以完全专注于业务功能

### eamples
在项目`examples`目录下，有一个简单的示例，可以接收网络请求，使用state store来储存和获取数据。
##### 1. 编译wasm代码
```shell
cd examples/service-call
make
```
##### 2. 运行dapr
安装并初始化dapr

在主目录下运行`make run-example`命令
##### 3. 发送请求
```shell
# 设置value
curl http://localhost:3500/v1.0/invoke/go-wasm-app/method/service-call/set -d 'value1'
# 查询，这里应该会返回value1
curl http://localhost:3500/v1.0/invoke/go-wasm-app/method/service-call/get
```

### 下一步
完善功能和细节
增加wasm instance pool，可以复用wasm instance以及对请求进行并行处理
现在是wasm->wasm-dapr(wasm host + dapr sdk) -> dapr,后面会探索直接长在dapr上，wasm->dapr.