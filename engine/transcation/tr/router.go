package tr

import (
	"marketing/delivery"
	"marketing/engine/transcation/model"
	tskM "marketing/task/model"

	"marketing/consts/engine"
)

func NewTr(task *tskM.Task) model.T {
	switch task.Type {
	case engine.Tr_HTTP:
		return &httpTr{Task: task}
	case engine.Tr_RPC:
		return GetRPCTask(task.Handler)
	case engine.Tr_Local:
		switch task.ID {
		case engine.TaskId_DeductQuota:
			return &deductQuota{}
		case engine.TaskId_Reward:
			return delivery.NewReward()
		default:
			return nil
		}
	default:
		return nil
	}
}
