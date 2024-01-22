package errs

var (
	Internal      = newErr(-1, "Internal error")
	Bind          = newErr(-2, "Biend error")
	InvalidParams = newErr(-3, "Invalid parameters")
	EmptyUsername = newErr(-4, "Username not found")
	NoAuthRes     = newErr(-5, "No authentication resource available")
	NoAuth        = newErr(-5, "Not granted")
)
