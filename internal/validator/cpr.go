package validator

import (
	"errors"
	"math"
	"strconv"
	"strings"
)

func ValidateCPR(cpr string) (bool, error) {
	cpr = strings.TrimSpace(cpr)
	idLen := len(cpr)

	if idLen == 0 {
		return false, errors.New("CPR is required")
	}
	if idLen != 9 {
		return false, errors.New("CPR must be 9 characters long")
	}
	num := 0
	// Loop through all characters except the last one
	for i := 0; i < idLen-1; i++ {
		// Math.Abs(i - id.Count())
		weight := int(math.Abs(float64(i - idLen)))
		if cpr[i] < '0' || cpr[i] > '9' {
			return false, errors.New("CPR must contain only digits")
		}
		asciiValue := int(cpr[i])
		num += weight * asciiValue
	}

	var x int
	z := num % 11
	if z == 0 || z == 1 {
		x = 0
	} else {
		x = 11 - z
	}

	// Compare the last character of the ID to the first character of string 'x'
	lastChar := cpr[idLen-1]
	xStrFirstChar := strconv.Itoa(x)[0]

	if lastChar == xStrFirstChar {
		return true, nil
	}
	return false, errors.New("invalid CPR checksum")

}
