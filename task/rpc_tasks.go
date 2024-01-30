package task

type RpcTask interface{}

var route = map[string]T{}

func GetRPCTask(handler string) T {
	return route[handler]
}

func RegisterRPCTask(handler string, task T) {
	route[handler] = task
}
