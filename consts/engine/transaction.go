package engine

type TaskType int

const (
	TaskType_HTTP TaskType = iota
	TaskType_RPC
)
