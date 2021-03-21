package main

import "fmt"

func InitCountLamp() int {
	var angka int
	fmt.Print("Masukkan angka : ")
	fmt.Scanln(&angka)
	return angka
}
func CountLamp() int {
	val := InitCountLamp()
	res := 0
	for i := 1; i < val; i++ {
		res = res + val/i
	}
	return res
}
