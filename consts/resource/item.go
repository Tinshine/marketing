package resource

type ItemType int

const (
	Currency ItemType = iota // virtual currency
	Credit
	CDKey
)
