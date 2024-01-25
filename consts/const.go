package consts

type ReleaseState int

const (
	StateCreated ReleaseState = iota
	StateRelasing
	StateOnline
	StateOffline
)
