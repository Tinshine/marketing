package resource

import (
	"marketing/consts/errs"

	"github.com/pkg/errors"
)

type ItemType int

const (
	Currency ItemType = iota // virtual currency
	Credit
	CDKey
	UnknownItemType
)

func (t ItemType) Validate() error {
	if t < Currency || t >= UnknownItemType {
		return errors.WithMessage(errs.InvalidParams, "unsupported item type")
	}
	return nil
}

type DeliveryType int

const (
	Sync DeliveryType = iota
	Async
)

func (d DeliveryType) Validate() error {
	if d != Sync && d != Async {
		return errors.WithMessage(errs.InvalidParams, "unsupported delivery type")
	}
	return nil
}
