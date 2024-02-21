package reward

import (
	"context"
	"marketing/consts"
	"marketing/consts/errs"
	"marketing/delivery/client"
	dM "marketing/delivery/model"
	"marketing/delivery/router"
	trM "marketing/engine/transcation/model"
	gM "marketing/manager/resource/model/gift"
	itemM "marketing/manager/resource/model/item"
	gServ "marketing/manager/resource/service/gift"
	itemServ "marketing/manager/resource/service/item"
	"marketing/quota"
	"marketing/util/log"
	"sync"

	"github.com/pkg/errors"
)

func parseParams(txId string, params *trM.Params) (*dM.RewardReq, error) {
	req := new(dM.RewardReq)
	qid, ok := params.Input["quota_id"]
	if !ok {
		return nil, errors.WithMessage(errs.InvalidParams, "quota_id is required")
	}
	req.QuotaId = qid.(uint)
	gid, ok := params.Input["group_id"]
	if !ok {
		return nil, errors.WithMessage(errs.InvalidParams, "group_id is required")
	}
	req.GroupId = gid.(uint)
	req.AppId = params.AppId
	req.Ev = params.Ev
	req.UserId = params.User.GetId()
	req.TxId = txId
	return req, nil
}

func tryOrder(ctx context.Context, req *dM.RewardReq) error {
	order, err := dM.FindOrder(ctx, req)
	if err != nil && err != errs.OrderNotFound {
		return errors.WithMessage(err, "find order")
	}
	if err == nil {
		// a duplicate order exists, should ignore this in case cancel or confirm has been performed
		log.Error("Try.FindOrder.Exist", errs.DuplicatedTry, "req", req, "order", order)
		return nil
	}
	// create a new order and update user's limit
	order, err = dM.CreateOrder(ctx, req, consts.StateTry)
	if err != nil {
		return errors.WithMessage(err, "create order")
	}
	log.Info("Try.CreateOrder.Success", "req", req, "order", order)

	if err := quota.TryDeductQuota(ctx, quota.NewDeductQuotaReq(
		req.UserId, req.QuotaId, req.TxId,
	)); err != nil {
		return errors.WithMessage(err, "deduct quota")
	}
	log.Info("Try.DeductQuota.Success", "req", req, "order", order)
	return nil
}

func cancelOrder(ctx context.Context, req *dM.RewardReq) error {
	order, err := dM.FindOrder(ctx, req)
	if err != nil && err != errs.OrderNotFound {
		return errors.WithMessage(err, "find order")
	}
	if err == errs.OrderNotFound {
		// empty cancel problem
		log.Error("Cancel.Empty.TxRecord", errs.TransactionEmptyCancel, "req", req)
		_, err = dM.CreateOrder(ctx, req, consts.StateCancel)
		if err != nil {
			return errors.WithMessage(err, "create cancel order")
		}
		log.Info("Cancel.Empty.CreateCancel", "req", req, "order", order)
		return nil
	}

	if err := dM.UpdateOrder(ctx,
		order.Id, req.Ev, consts.StateTry, consts.StateCancel); err != nil {
		return errors.WithMessage(err, "update order")
	}
	return nil
}

func confirmOrder(ctx context.Context, req *dM.RewardReq) error {
	order, err := dM.FindOrder(ctx, req)
	if err != nil && err != errs.OrderNotFound {
		return errors.WithMessage(err, "find order")
	}
	if err == errs.OrderNotFound {
		// empty confirm problem
		log.Error("Confirm.Empty.TxRecord", errs.TransactionEmptyConfirm, "req", req)
		_, err = dM.CreateOrder(ctx, req, consts.StateConfirm)
		if err != nil {
			return errors.WithMessage(err, "create confirm order")
		}
		log.Info("Confirm.Empty.CreateConfirm", "req", req, "order", order)
		return nil
	}

	if err := dM.UpdateOrder(ctx,
		order.Id, req.Ev, consts.StateTry, consts.StateConfirm); err != nil {
		return errors.WithMessage(err, "update order")
	}
	return nil
}

func delivery(ctx context.Context, req *dM.RewardReq) error {
	gift, err := getGift(ctx, req)
	if err != nil {
		return errors.WithMessage(err, "get gift")
	}

	clis := make([]client.D, 0, len(gift.Items))
	for i := range gift.Items {
		item, err := getItem(ctx, req.AppId, gift.Items[i].ItemId)
		if err != nil {
			return errors.WithMessage(err, "get item")
		}
		cli, err := router.GetClient(item.ItemType)
		if err != nil {
			return errors.WithMessage(errs.UnsupportedItemType, "unsupported item type")
		}
		clis = append(clis, cli)
	}

	wg := &sync.WaitGroup{}
	wg.Add(len(clis))
	for i := 0; i < len(clis); i++ {
		go func(i int) {
			defer wg.Done()
			r := &client.DeliveryReq{}
			resp, err := clis[i].Delivery(ctx, r)
			if err != nil {
				log.Error("delivery.Delivery.err", err, "req", r, "gift", gift, "item", gift.Items[i])
			}
			log.Info("delivery.Delivery.Success", "req", r, "gift", gift,
				"item", gift.Items[i], "resp", resp)
		}(i)
	}
	wg.Wait()
	return nil
}

func getGift(ctx context.Context, req *dM.RewardReq) (*gM.RespModel, error) {
	gResp, err := gServ.Query(ctx, &gM.QueryReq{
		AppId:   req.AppId,
		GroupId: &req.GroupId,
		Env:     req.Ev,
	})
	if err != nil {
		return nil, errors.WithMessage(err, "query gift group")
	}
	if len(gResp.Data) == 0 {
		return nil, errors.WithMessage(errs.GiftGroupNotFound, "gift group not found")
	}
	if len(gResp.Data) > 1 {
		return nil, errors.WithMessage(errs.Internal, "gift group not unique")
	}
	return gResp.Data[0], err
}

func getItem(ctx context.Context, appId uint, itemId int) (*itemM.RespModel, error) {
	resp, err := itemServ.Query(ctx, &itemM.QueryReq{
		AppId:  appId,
		ItemId: &itemId,
	})
	if err != nil {
		return nil, errors.WithMessage(err, "query item")
	}
	if len(resp.Data) == 0 {
		return nil, errors.WithMessage(errs.ItemNotFound, "item not found")
	}
	if len(resp.Data) > 1 {
		return nil, errors.WithMessage(errs.Internal, "item not unique")
	}
	return resp.Data[0], err
}
