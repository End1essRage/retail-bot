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

	CategoryPrefix = "c"
	ProductPrefix  = "p"
	BackPrefix     = "b"
)
