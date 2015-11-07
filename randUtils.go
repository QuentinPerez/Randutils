package randutils

import (
	"crypto/rand"
	"errors"
)

func getBuffer(length uint32) []byte {
	buff := make([]byte, length)

	rand.Read(buff)
	return buff
}

// AlphaLower returns an array
// which contains characters between:
//     'a' and 'z'
func AlphaLower(length uint32) ([]byte, error) {
	return Pattern(length, "abcdefghijklmnopqrstuvwxyz")
}

// AlphaUpper returns an array
// which contains characters between:
//     'A' and 'Z'
func AlphaUpper(length uint32) ([]byte, error) {
	return Pattern(length, "ABCDEFGHIJKLMNOPQRSTUVWXYZ")
}

// AlphaUpperNum returns an array
// which contains characters between:
//     'A' and 'Z'
//     '0' and '9'
func AlphaUpperNum(length uint32) ([]byte, error) {
	return Pattern(length, "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
}

// AlphaLowerNum returns an array
// which contains characters between:
//     'a' and 'z'
//     '0' and '9'
func AlphaLowerNum(length uint32) ([]byte, error) {
	return Pattern(length, "abcdefghijklmnopqrstuvwxyz0123456789")
}

// Printable returns an array
// which contains characters printable:
func Printable(length uint32) ([]byte, error) {
	return Pattern(length, `!#$%&()*+,-./0123456789:;<=>?@abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ[\]^_{|}`+"`")
}

// Numeric returns an array
// which contains characters between:
//     '0' and '9'
func Numeric(length uint32) ([]byte, error) {
	return Pattern(length, "0123456789")
}

func Pattern(length uint32, pattern string) ([]byte, error) {
	lengthPattern := len(pattern)

	if lengthPattern == 0 {
		return nil, errors.New("The pattern's length can not be 0")
	}
	buff := getBuffer(length)
	for i, c := range buff {
		buff[i] = []byte(pattern)[int(c)%lengthPattern]
	}
	return buff, nil
}
