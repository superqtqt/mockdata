package cmd

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var Random = func() float64 {
	return rand.Float64()
}
