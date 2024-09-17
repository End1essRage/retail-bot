package constants

const (
	ENV_LOCAL = "LOCAL"
	ENV_DEV   = "DEV"
)

type CallBackType string
type CacheType string

const (
	//buttonns
	//menu
	ProductSelect  CallBackType = "p"
	Back           CallBackType = "b"
	CategorySelect CallBackType = "c"
	ProductAdd     CallBackType = "pa"

	//cart
	ProductIncrement CallBackType = "pi"
	ProductDecrement CallBackType = "pd"
	ProductAmount    CallBackType = "pr"
	ProductName      CallBackType = "pn"

	ClearCart   CallBackType = "cc"
	CreateOrder CallBackType = "co"

	//cache
	CacheCartUserPrefix CacheType = "cu"
	CacheSeparator                = "_"
	MenuKey                       = "menu"

	//buttons
	TypeSeparator = "_"
	DataSeparator = "|"
	FlagSeparator = "="
)
