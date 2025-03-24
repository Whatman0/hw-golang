package request

import (
	"math/rand"
	"strconv"
	"time"
)

func HashIt() string {
	rand.NewSource(time.Now().UnixNano())
	num := rand.Intn(999999999)
	return strconv.Itoa(num)
}
