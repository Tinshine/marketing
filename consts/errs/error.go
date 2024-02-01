package errs

var (
	Internal              = newErr(-1, "Internal error")
	Bind                  = newErr(-2, "Bind error")
	InvalidParams         = newErr(-3, "Invalid parameters")
	EmptyUsername         = newErr(-4, "Username not found")
	NoAuthRes             = newErr(-5, "No authentication resource available")
	NoAuth                = newErr(-5, "Not granted")
	ItemNotFound          = newErr(-6, "Item not found")
	DuplicateItem         = newErr(-7, "Duplicate item")
	NoDeliveryClient      = newErr(-8, "No delivery client available")
	ConfNotFound          = newErr(-9, "Conf not found")
	COnfEnvNotExist       = newErr(-10, "Conf env not exist")
	RedisException        = newErr(-11, "Unexpected redis error")
	UnlockOthers          = newErr(-12, "Unlock others redis lock")
	TooManyRequests       = newErr(-13, "Too many requests")
	DeleteNotAllowed      = newErr(-14, "Delete not allowed")
	InvalidState          = newErr(-15, "Invalid state")
	DuplicateRule         = newErr(-16, "Duplicate rule")
	QuotaLimit            = newErr(-17, "Quota limit exceeded")
	ConstraintsNotMeet    = newErr(-18, "Constraints not meet")
	GormScanWrongType     = newErr(-19, "Gorm scan wrong type")
	BatchQueryLenNotMatch = newErr(-20, "Batch query len not match")
	TransactionConfirmed  = newErr(-21, "Transaction already confirmed")
	TransactionCanceled   = newErr(-22, "Transaction already canceled")
	DuplicateRpcTask      = newErr(-23, "Duplicate rpc task")
)
