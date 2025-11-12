package middleware

import (
	"base/data"

	"github.com/m-goku/rkt"
)

type Middleware struct {
	App    *rkt.RKT
	Models data.Models
}
