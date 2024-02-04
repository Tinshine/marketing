package quota

type DeductQuotaReq struct {
	UserId  string
	QuotaId uint
	TxId    uint
}

func NewDeductQuotaReq(userId string, quotaId uint, txId uint) *DeductQuotaReq {
	return &DeductQuotaReq{
		UserId:  userId,
		QuotaId: quotaId,
		TxId:    txId,
	}
}
