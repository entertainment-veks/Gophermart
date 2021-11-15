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

func stringToInt(input string) (int, error) {
	return strconv.Atoi(input)
}

func IsValid(input string) bool {
	inted, err := stringToInt(input)
	if err != nil {
		return false
	}
	return luhn.Valid(inted)
}
