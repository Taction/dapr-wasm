/**
 * @Author: zhangchao
 * @Description:
 * @Date: 2023/3/4 5:37 PM
 */
package runtime

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"regexp"

	"github.com/knqyf263/go-plugin/types/known/anypb"

	"github.com/taction/dapr-wasm/proto/runtime/v1"
)

type AppNameCtx struct {
}

type Option func(w http.ResponseWriter)

func (rt *Runtime) IndexHandler(w http.ResponseWriter, r *http.Request) {
	isFound := false
	for p, app := range rt.appMap {
		// 这里循环匹配Path，先添加的先匹配
		reg, err := regexp.Compile(p)
		if err != nil {
			continue
		}
		if reg.MatchString(r.URL.Path) {
			isFound = true
			ctx := context.Background()
			ctx = context.WithValue(ctx, AppNameCtx{}, r)
			body, err := io.ReadAll(r.Body)
			if err != nil {
				respond(w, withError(err))
				return
			}
			res, err := app.OnInvoke(ctx, runtime.InvokeRequest{Method: r.URL.Path, Data: &anypb.Any{Value: body}})
			if err != nil {
				respond(w, withError(err))
				return
			}
			respond(w, withInvokeResponse(res))
		}
	}
	if !isFound {
		// 未匹配到路由
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "404 Page Not Found!")
	}
}
func (rt *Runtime) HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "OK")
}

func withInvokeResponse(res runtime.InvokeResponse) Option {
	return func(w http.ResponseWriter) {
		if len(res.GetContentType()) > 0 {
			w.Header().Set("Content-Type", res.GetContentType())
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, string(res.GetData().GetValue()))
	}
}

func withError(err error) Option {
	return func(w http.ResponseWriter) {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
	}
}

func respond(w http.ResponseWriter, option ...Option) {
	for _, o := range option {
		o(w)
	}
}
