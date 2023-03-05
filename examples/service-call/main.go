//go:build tinygo.wasm

package main

import (
	"context"
	"errors"
	"strings"

	"github.com/knqyf263/go-plugin/types/known/anypb"
	"github.com/knqyf263/go-plugin/types/known/emptypb"

	"github.com/taction/dapr-wasm/proto/runtime/v1"
)

// main is required for TinyGo to compile to Wasm.
func main() {
	runtime.RegisterAppCallback(ServiceCall{})
}

const StoreName = "statestore"
const StoreKey = "key"

type ServiceCall struct{}

func (g ServiceCall) OnInvoke(ctx context.Context, request runtime.InvokeRequest) (runtime.InvokeResponse, error) {
	if strings.HasSuffix(request.Method, "/set") {
		hostFunctions := runtime.NewDapr()
		_, err := hostFunctions.SaveState(ctx, runtime.SaveStateRequest{StoreName: StoreName, States: []*runtime.StateItem{{Key: StoreKey, Value: request.GetData().GetValue()}}})
		if err != nil {
			return runtime.InvokeResponse{}, err
		}
		return runtime.InvokeResponse{Data: &anypb.Any{Value: []byte("OK")}}, nil
	} else if strings.HasSuffix(request.Method, "/get") {
		hostFunctions := runtime.NewDapr()
		res, err := hostFunctions.GetState(ctx, runtime.GetStateRequest{StoreName: StoreName, Key: StoreKey})
		if err != nil {
			return runtime.InvokeResponse{}, err
		}
		return runtime.InvokeResponse{Data: &anypb.Any{Value: []byte(res.GetData())}}, nil
	} else {
		return runtime.InvokeResponse{Data: &anypb.Any{Value: []byte("unknown method" + request.GetMethod())}}, errors.New("unknown method")
	}
}

func (g ServiceCall) ListTopicSubscriptions(ctx context.Context, empty emptypb.Empty) (runtime.ListTopicSubscriptionsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (g ServiceCall) OnTopicEvent(ctx context.Context, request runtime.TopicEventRequest) (runtime.TopicEventResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (g ServiceCall) ListInputBindings(ctx context.Context, empty emptypb.Empty) (runtime.ListInputBindingsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (g ServiceCall) OnBindingEvent(ctx context.Context, request runtime.BindingEventRequest) (runtime.BindingEventResponse, error) {
	//TODO implement me
	panic("implement me")
}
