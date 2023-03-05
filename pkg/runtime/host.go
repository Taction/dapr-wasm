/**
 * @Author: zhangchao
 * @Description:
 * @Date: 2023/3/4 5:01 PM
 */
package runtime

import (
	"context"

	"github.com/knqyf263/go-plugin/types/known/emptypb"

	"github.com/taction/dapr-wasm/proto/runtime/v1"
)

var _ runtime.Dapr = (*Runtime)(nil)

//type Host struct {
//
//}

func (rt *Runtime) InvokeService(ctx context.Context, request runtime.InvokeServiceRequest) (runtime.InvokeResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (rt *Runtime) GetState(ctx context.Context, request runtime.GetStateRequest) (runtime.GetStateResponse, error) {
	res, err := rt.daprClient.GetState(ctx, request.StoreName, request.Key, nil)
	if err != nil {
		return runtime.GetStateResponse{}, err
	}
	return runtime.GetStateResponse{Data: res.Value}, nil
}

func (rt *Runtime) GetBulkState(ctx context.Context, request runtime.GetBulkStateRequest) (runtime.GetBulkStateResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (rt *Runtime) SaveState(ctx context.Context, request runtime.SaveStateRequest) (emptypb.Empty, error) {
	err := rt.daprClient.SaveState(ctx, request.StoreName, request.GetStates()[0].GetKey(), request.GetStates()[0].GetValue(), nil)
	if err != nil {
		return emptypb.Empty{}, err
	}
	return emptypb.Empty{}, nil
}

func (rt *Runtime) QueryStateAlpha1(ctx context.Context, request runtime.QueryStateRequest) (runtime.QueryStateResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (rt *Runtime) DeleteState(ctx context.Context, request runtime.DeleteStateRequest) (emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (rt *Runtime) DeleteBulkState(ctx context.Context, request runtime.DeleteBulkStateRequest) (emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (rt *Runtime) ExecuteStateTransaction(ctx context.Context, request runtime.ExecuteStateTransactionRequest) (emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (rt *Runtime) PublishEvent(ctx context.Context, request runtime.PublishEventRequest) (emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (rt *Runtime) BulkPublishEventAlpha1(ctx context.Context, request runtime.BulkPublishRequest) (runtime.BulkPublishResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (rt *Runtime) InvokeBinding(ctx context.Context, request runtime.InvokeBindingRequest) (runtime.InvokeBindingResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (rt *Runtime) GetSecret(ctx context.Context, request runtime.GetSecretRequest) (runtime.GetSecretResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (rt *Runtime) GetBulkSecret(ctx context.Context, request runtime.GetBulkSecretRequest) (runtime.GetBulkSecretResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (rt *Runtime) RegisterActorTimer(ctx context.Context, request runtime.RegisterActorTimerRequest) (emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (rt *Runtime) UnregisterActorTimer(ctx context.Context, request runtime.UnregisterActorTimerRequest) (emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (rt *Runtime) RegisterActorReminder(ctx context.Context, request runtime.RegisterActorReminderRequest) (emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (rt *Runtime) UnregisterActorReminder(ctx context.Context, request runtime.UnregisterActorReminderRequest) (emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (rt *Runtime) RenameActorReminder(ctx context.Context, request runtime.RenameActorReminderRequest) (emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (rt *Runtime) GetActorState(ctx context.Context, request runtime.GetActorStateRequest) (runtime.GetActorStateResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (rt *Runtime) ExecuteActorStateTransaction(ctx context.Context, request runtime.ExecuteActorStateTransactionRequest) (emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (rt *Runtime) InvokeActor(ctx context.Context, request runtime.InvokeActorRequest) (runtime.InvokeActorResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (rt *Runtime) GetConfigurationAlpha1(ctx context.Context, request runtime.GetConfigurationRequest) (runtime.GetConfigurationResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (rt *Runtime) SubscribeConfigurationAlpha1(ctx context.Context, request runtime.SubscribeConfigurationRequest) (runtime.SubscribeConfigurationResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (rt *Runtime) UnsubscribeConfigurationAlpha1(ctx context.Context, request runtime.UnsubscribeConfigurationRequest) (runtime.UnsubscribeConfigurationResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (rt *Runtime) TryLockAlpha1(ctx context.Context, request runtime.TryLockRequest) (runtime.TryLockResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (rt *Runtime) UnlockAlpha1(ctx context.Context, request runtime.UnlockRequest) (runtime.UnlockResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (rt *Runtime) GetMetadata(ctx context.Context, empty emptypb.Empty) (runtime.GetMetadataResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (rt *Runtime) SetMetadata(ctx context.Context, request runtime.SetMetadataRequest) (emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (rt *Runtime) StartWorkflowAlpha1(ctx context.Context, request runtime.StartWorkflowRequest) (runtime.WorkflowReference, error) {
	//TODO implement me
	panic("implement me")
}

func (rt *Runtime) GetWorkflowAlpha1(ctx context.Context, request runtime.GetWorkflowRequest) (runtime.GetWorkflowResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (rt *Runtime) TerminateWorkflowAlpha1(ctx context.Context, request runtime.TerminateWorkflowRequest) (runtime.TerminateWorkflowResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (rt *Runtime) Shutdown(ctx context.Context, empty emptypb.Empty) (emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}
