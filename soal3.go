package main

import (
	"fmt"
	"strconv"
	"strings"
)

func InitDetailsNumber() int {
	var angka int
	fmt.Print("Masukkan angka : ")
	fmt.Scanln(&angka)
	return angka
}
func DetailsNumber() {
	val := InitDetailsNumber()
	val_str := strconv.Itoa(val)
	length := len(val_str) - 1
	for i := 0; i <= length; i++ {
		fmt.Println(string(val_str[i]) + strings.Repeat("0", length-i))
	}
}
