package constants

const (
	ENV_LOCAL = "LOCAL"
	ENV_DEV   = "DEV"
)

type CallBackType int64

const (
	ProductSelect CallBackType = 0
	Back          CallBackType = 1
	//categories
	CategorySelect CallBackType = 2
	ProductAdd     CallBackType = 3

	CategoryPrefix         = "c"
	ProductPrefix          = "p"
	ProductAddPrefix       = "pa"
	ProductIncrementPrefix = "pi"
	ProductDecrementPrefix = "pd"
	ProductAmountPrefix    = "pr"
	ProductNamePrefix      = "pn"
	BackPrefix             = "b"

	CacheCartUserPrefix = "cu"
)
