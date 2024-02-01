package tr

import (
	"marketing/consts/errs"
	"marketing/engine/transcation/model"
)

var route = map[string]model.T{}

func GetRPCTask(handler string) model.T {
	return route[handler]
}

func RegisterRPCTask(handler string, tr model.T) error {
	if _, exist := route[handler]; exist {
		return errs.DuplicateRpcTask
	}
	route[handler] = tr
	return nil
}
