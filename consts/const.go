package consts

type ReleaseState int

const (
	StateCreated ReleaseState = iota
	StateRelasing
	StateOnline
	StateOffline
)

type Env int

const (
	Dev Env = iota
	Test
	Prod
)

type LoginType int

const (
	Sdk LoginType = iota
	Passport
	Web
)

type TxState int

const (
	StateTry TxState = iota
	StateCancel
	StateConfirm
	StateRetry
)
