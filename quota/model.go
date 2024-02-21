package quota

type DeductQuotaReq struct {
	UserId  string
	QuotaId uint
	TxId    string
}

func NewDeductQuotaReq(userId string, quotaId uint, txId string) *DeductQuotaReq {
	return &DeductQuotaReq{
		UserId:  userId,
		QuotaId: quotaId,
		TxId:    txId,
	}
}
