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
	ProductSelect  CallBackType = "ps"
	Back           CallBackType = "b"
	CategorySelect CallBackType = "cs"

	//cart
	ProductAdd       CallBackType = "pa"
	ProductIncrement CallBackType = "pi"
	ProductDecrement CallBackType = "pd"
	ClearCart        CallBackType = "cc"
	CreateOrder      CallBackType = "co"
	ProductAmount    CallBackType = "pam" //не используется
	ProductName      CallBackType = "pn"  //не используется

	//order

	//cache
	CacheCartUserPrefix CacheType = "cu"
	CacheSeparator                = "_"
	MenuKey                       = "menu"

	//buttons
	TypeSeparator = "_"
	DataSeparator = "|"
	FlagSeparator = "="
)
