package main

import "fmt"

func InitSocks() int {
	var angka int
	fmt.Print("masukkan angka : ")
	fmt.Scanln(&angka)
	return angka
}
func Socks() int {
	val := InitSocks()
	res := 0
	arr := Input(val)
	u_arr := Unique(arr)
	for i := 0; i < len(u_arr); i++ {
		count := 0
		for j := 0; j < len(arr); j++ {
			if u_arr[i] == arr[j] {
				count += 1
			}
		}
		res = res + count/2
	}
	return res
}

func Input(val int) []int {
	var num int
	var res []int
	for i := 0; i < val; i++ {
		fmt.Print("input color number: ")
		fmt.Scanln(&num)
		res = append(res, num)
	}
	fmt.Println(res)
	return res
}

func Unique(arr []int) []int {
	occured := map[int]bool{}
	result := []int{}
	for e := range arr {
		if occured[arr[e]] != true {
			occured[arr[e]] = true
			result = append(result, arr[e])
		}
	}
	return result
}
