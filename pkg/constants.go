package constants

const (
	ENV_LOCAL = "LOCAL"
	ENV_DEV   = "DEV"
)

type CallBackType int64

const (
	//buttonns
	//menu
	CategoryPrefix   = "c"
	ProductPrefix    = "p"
	ProductAddPrefix = "pa"
	BackPrefix       = "b"

	ProductSelect  CallBackType = 0
	Back           CallBackType = 1
	CategorySelect CallBackType = 2
	ProductAdd     CallBackType = 3

	//cart
	ProductIncrementPrefix = "pi"
	ProductDecrementPrefix = "pd"
	ProductAmountPrefix    = "pr"
	ProductNamePrefix      = "pn"

	ProductIncrement CallBackType = 4
	ProductDecrement CallBackType = 5
	ProductAmount    CallBackType = 6
	ProductName      CallBackType = 7

	ClearCartPrefix   = "cc"
	CreateOrderPrefix = "co"

	ClearCart   CallBackType = 8
	CreateOrder CallBackType = 9

	//cache
	CacheCartUserPrefix = "cu"
)
