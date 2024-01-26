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
