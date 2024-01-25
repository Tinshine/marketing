package client

import (
	"context"
	"marketing/consts/resource"
)

type C interface {
	Delivery(context.Context, *DeliveryReq) (*DeliveryResp, error)
}

type DeliveryReq struct {
	AppId    uint
	ItemType resource.ItemType
	ItemId   uint
	Count    uint
	OrderId  string
}

type DeliveryResp struct {
}
