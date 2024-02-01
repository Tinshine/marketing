package engine

type TrType int

const (
	Tr_HTTP TrType = iota
	Tr_RPC
	Tr_Local
)

const (
	TaskId_DeductQuota uint = 1
	TaskId_Reward      uint = 2
)
