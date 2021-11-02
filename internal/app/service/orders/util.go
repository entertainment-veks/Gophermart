package orders

import (
	"strconv"

	"github.com/theplant/luhn"
)

const (
	StatusNew        = "NEW"
	StatusProcessing = "PROCESSING"
	StatusInvalid    = "INVALID"
	StatusProcessed  = "PROCESSED"
)

func bytesToInt(input []byte) (int, error) {
	stringed := string(input)
	return strconv.Atoi(stringed)
}

func isValid(input int) bool {
	return luhn.Valid(input)
}
