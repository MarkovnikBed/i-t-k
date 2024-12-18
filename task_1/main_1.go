package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
)

func main() {
	var numDecimal int = 42
	var numOctal int = 052
	var numHexadecimal int = 0x2A
	var pi float64 = 3.14
	var name string = "Golang"
	var isActive bool = true
	var complexNum complex64 = 1 + 2i

	fmt.Print(getTypes(numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum))
	str := concatinate(numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum)
	fmt.Print(getHash(str, "go-2024"))

}

func getHash(str string, salt string) string {
	// я конвертировал строку в слайс байт так как посчитал что в задании ошибка (требование преобразовать строку в слайс рун и получить хэш)
	strByte := append(append([]byte(str[:len(str)/2]), []byte(salt)...), []byte(str[len(str)/2:])...)
	hash := sha256.Sum256(strByte)
	return hex.EncodeToString(hash[:])
}

func getTypes(nums ...any) string {
	var builder strings.Builder
	for _, num := range nums {
		builder.WriteString(fmt.Sprintf("%v is %T type\n", num, num))
	}
	return builder.String()
}

func concatinate(nums ...any) string {
	var builder strings.Builder
	for _, v := range nums {
		builder.WriteString(fmt.Sprintf("%v", v))
	}
	return builder.String()
}
