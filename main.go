/**
 * @Author: zhangchao
 * @Description:
 * @Date: 2023/3/1 9:07 PM
 */
package main

import (
	"flag"
	"log"
	"strconv"

	"github.com/taction/dapr-wasm/pkg/runtime"
)

var (
	config   = flag.String("config", "config.yaml", "Runtime config")
	httpPort = flag.String("port", "6003", "HTTP port to listen on")
)

func main() {
	flag.Parse()
	port, _ := strconv.Atoi(*httpPort)
	rt := runtime.NewRuntime(&runtime.Config{
		Port:     port,
		WASMPath: *config,
	})
	err := rt.Run()
	if err != nil {
		log.Fatalf("failed to run runtime: %v", err)
	}
}
