package router

import (
	"marketing/consts/errs"
	"marketing/consts/resource"
	"marketing/delivery/client"
	"marketing/delivery/services/credit"
)

var rt = map[resource.ItemType]client.C{
	resource.Credit: &credit.Credit{},
}

func GetClient(itemType resource.ItemType) (client.C, error) {
	c, ok := rt[itemType]
	if !ok {
		return nil, errs.NoDeliveryClient
	}
	return c, nil
}
