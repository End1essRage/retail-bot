package constants

const (
	ENV_LOCAL = "LOCAL"
	ENV_DEV   = "DEV"
)

type CallBackType string
type CacheType string
type OrderStatus int

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
	OrderShortOpen  CallBackType = "oop"
	OrderBackToList CallBackType = "obb"
	OrderCancel     CallBackType = "oca"
	OrderClose      CallBackType = "ocl"
	OrderApply      CallBackType = "oap"

	//cache
	CacheCartUserPrefix CacheType = "cu"
	CacheSeparator                = "_"
	MenuKey                       = "menu"

	New       OrderStatus = 0
	Accepted  OrderStatus = 1
	Completed OrderStatus = 2
	Cancelled OrderStatus = 3
	//buttons
	TypeSeparator = "_"
	DataSeparator = "|"
	FlagSeparator = "="
)
