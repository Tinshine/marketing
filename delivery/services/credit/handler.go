package credit

import (
	"context"
	"marketing/delivery/client"
)

type Credit struct {
}

func (c *Credit) Delivery(ctx context.Context, rq *client.DeliveryReq) (*client.DeliveryResp, error) {
	// todo...
	return nil, nil
}
