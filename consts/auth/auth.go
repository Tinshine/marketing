package auth

type FailReason string

const (
	NotGranted FailReason = "not granted"
	Expired    FailReason = "expired"
)
