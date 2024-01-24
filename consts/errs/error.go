package errs

var (
	Internal      = newErr(-1, "Internal error")
	Bind          = newErr(-2, "Bind error")
	InvalidParams = newErr(-3, "Invalid parameters")
	EmptyUsername = newErr(-4, "Username not found")
	NoAuthRes     = newErr(-5, "No authentication resource available")
	NoAuth        = newErr(-5, "Not granted")
	ItemNotFound  = newErr(-6, "Item not found")
	DuplicateItem = newErr(-7, "Duplicate item")
)
