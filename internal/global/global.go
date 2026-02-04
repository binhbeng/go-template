package global

import "time"

const (
	Version = "0.3.0"
	PerPage = 10

	RedisPrefix           = "goex:"
	DefaultRequestTimeout = 5 * time.Second

	Issuer  = "go-ex"
)
