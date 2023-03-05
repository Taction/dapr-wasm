/**
 * @Author: zhangchao
 * @Description:
 * @Date: 2023/3/2 8:38 PM
 */
package runtime

import (
	"errors"
	"log"
	"os"
	"sync"
	"time"

	dapr "github.com/dapr/go-sdk/client"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"

	"github.com/taction/dapr-wasm/pkg/config"
	"github.com/taction/dapr-wasm/pkg/utils"
)

var (
	defaultTimeout = 30 * time.Second
)

type Runtime struct {
	lock       sync.Mutex
	daprClient dapr.Client
	config     *Config
}

type Application struct {
	Path string
}

func NewRuntime(c *Config) *Runtime {
	return &Runtime{}
}

func (r *Runtime) Run() error {
	if err := r.initClient(); err != nil {
		return err
	}

	return nil
}

func (r *Runtime) initClient() error {
	r.lock.Lock()
	defer r.lock.Unlock()
	timeout := defaultTimeout
	timer := time.After(timeout)
	for {
		select {
		case <-timer:
			return errors.New("connect timeout")
		default:
			client, err := dapr.NewClient()
			if err == nil {
				r.daprClient = client
				return nil
			}
			time.Sleep(1 * time.Second)
		}
	}
}

func (r *Runtime) initFromConfig() error {
	c, err := loadConfigFromFile(r.config.WASMPath)
	if err != nil {
		return err
	}
	for _, app := range c.Applications {
		if app.Trigger.Type == "wasm" {

		}
	}
	return nil
}

func loadConfigFromFile(path string) (*config.Config, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		logrus.Warnf("load config error when reading file %s : %s", path, err)
		return nil, err
	}
	c := &config.Config{}
	err = yaml.Unmarshal(b, c)
	return c, err
}

func (r *Runtime) run() error {
	// create the client
	client, err := dapr.NewClient()
	if err != nil {
		panic(err)
	}

	log.Printf("Listening on http://localhost:%d", appPort)
	utils.StartServer(appPort, appRouter, true, false)
}
