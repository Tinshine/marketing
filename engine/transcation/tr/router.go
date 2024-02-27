package tr

import (
	"marketing/delivery/reward"
	"marketing/engine/transcation/model"
	tskM "marketing/task/model"

	"marketing/consts/engine"
)

func NewTr(task *tskM.Task, txId string) model.T {
	switch task.Type {
	case engine.Tr_HTTP:
		return &httpTr{Task: task}
	case engine.Tr_RPC:
		return GetRPCTask(task.Handler)
	case engine.Tr_Local:
		switch task.Id {
		case engine.TaskId_DeductQuota:
			return &deductQuota{}
		case engine.TaskId_Reward:
			return reward.NewReward(txId)
		default:
			return nil
		}
	default:
		return nil
	}
}
