package constants

const (
	ENV_LOCAL = "LOCAL"
	ENV_DEV   = "DEV"
)

type CallBackType string
type CacheType string
type UserRole int
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
	//status
	OrderChangeStatus CallBackType = "ocs"
	OrderCancel       CallBackType = "oca"
	OrderClose        CallBackType = "ocl"
	OrderAccept       CallBackType = "oap"

	//cache
	CacheCartUserPrefix    CacheType = "cu"
	CacheProductNamePrefix CacheType = "pn"
	CacheSeparator                   = "_"
	MenuKey                          = "menu"

	New       OrderStatus = 0
	Accepted  OrderStatus = 1
	Completed OrderStatus = 2
	Cancelled OrderStatus = 3
	Closed    OrderStatus = 4
	//buttons
	TypeSeparator = "_"
	DataSeparator = "|"
	FlagSeparator = "="

	//roles
	Admin   UserRole = 0
	Manager UserRole = 1
	Client  UserRole = 2
)
