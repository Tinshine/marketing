package resource

type GiftType int

const (
	Normal GiftType = iota // normal gift
	Lottery
)

const (
	RedisPrefixSync = "lock_sync"
)
