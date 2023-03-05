/**
 * @Author: zhangchao
 * @Description:
 * @Date: 2023/3/2 8:38 PM
 */
package runtime

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	dapr "github.com/dapr/go-sdk/client"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"

	"github.com/taction/dapr-wasm/pkg/config"
	"github.com/taction/dapr-wasm/pkg/utils"
	"github.com/taction/dapr-wasm/proto/runtime/v1"
)

var (
	defaultTimeout = 30 * time.Second
)

type Runtime struct {
	lock       sync.Mutex
	daprClient dapr.Client
	config     *Config
	subs       []*subscription
	appMap     map[string]runtime.AppCallback
}

type Application struct {
	Path string
}

func NewRuntime(c *Config) *Runtime {
	return &Runtime{config: c, appMap: make(map[string]runtime.AppCallback)}
}

func (rt *Runtime) Run() error {
	if err := rt.initClient(); err != nil {
		return err
	}
	if err := rt.initFromConfig(); err != nil {
		return err
	}
	err := rt.run()
	if err != nil {
		return err
	}

	return nil
}

func (rt *Runtime) initClient() error {
	rt.lock.Lock()
	defer rt.lock.Unlock()
	timeout := defaultTimeout
	timer := time.After(timeout)
	for {
		select {
		case <-timer:
			return errors.New("connect timeout")
		default:
			client, err := dapr.NewClient()
			if err == nil {
				rt.daprClient = client
				return nil
			}
			time.Sleep(1 * time.Second)
		}
	}
}

func (rt *Runtime) initFromConfig() error {
	c, err := loadConfigFromFile(rt.config.WASMPath)
	if err != nil {
		return err
	}
	for _, app := range c.Applications {
		if app.Source.Type != "file" {
			return errors.New("unsupported source type: " + app.Source.Type)
		}
		if app.Trigger.Type == config.TriggerTypeHTTP {
			ctx := context.Background()
			p, err := runtime.NewAppCallbackPlugin(ctx, runtime.AppCallbackPluginOption{})
			if err != nil {
				return err
			}
			ins, err := p.Load(ctx, app.Source.Path, rt)
			if err != nil {
				return err
			}
			rt.appMap[app.Name] = ins
		} else if app.Trigger.Type == config.TriggerTypePubSub {
			return errors.New("not implemented yet")
		} else if app.Trigger.Type == config.TriggerTypeBinding {
			return errors.New("not implemented yet")
		} else {
			return errors.New("invalid trigger type")
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

func (rt *Runtime) run() error {
	errch := make(chan error)
	log.Printf("Listening on http://localhost:%d", rt.config.Port)
	utils.StartServer(rt.config.Port, rt.appRouter(), true, false, errch)
	err, ok := <-errch
	if ok {
		return err
	}
	return nil
}

// appRouter initializes restful api router
func (rt *Runtime) appRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	//router.HandleFunc("/*", rt.IndexHandler).Methods("GET", "POST", "OPTION", "DELETE")
	router.PathPrefix("/").Handler(http.HandlerFunc(rt.IndexHandler))
	router.HandleFunc("/healthz", rt.HealthHandler).Methods("GET", "POST")
	router.Use(mux.CORSMethodMiddleware(router))

	return router
}
