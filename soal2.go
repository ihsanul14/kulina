package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func InitMountainorValley() ([]string, int) {
	var (
		mode  int
		angka int
		res   []string
	)
	reader := bufio.NewReader(os.Stdin)

	// input hitung "Mountain" atau "Valley"
	fmt.Print("Masukkan angka: ")
	fmt.Scanln(&angka)
	fmt.Print("Masukkan Valley(0) atau Mountain(1): ")
	fmt.Scanln(&mode)
	for i := 0; i < angka; i++ {
		fmt.Print("Masukkan huruf U atau D :")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		res = append(res, strings.ToUpper(string(text[0])))
	}
	fmt.Println(res)
	return res, mode
}
func Step() string {
	var keterangan string
	val, mode := InitMountainorValley()
	step := 0
	position := 0
	prev := 0
	for i := 0; i < len(val); i++ {
		prev = step
		if val[i] == "D" {
			step -= 1
		} else if val[i] == "U" {
			step += 1
		} else {
			fmt.Println("invalid")
			break
		}
		if mode == 0 {
			position += Valley(prev, step)
			keterangan = "valley"
		} else if mode == 1 {
			position += Mountain(prev, step)
			keterangan = "mountain"
		}
	}
	result := fmt.Sprintf("Jumlah %s adalah %d", keterangan, position)
	return result
}

func Valley(prev int, next int) int {
	res := 0
	if next == 0 && prev < 0 {
		res += 1
	}
	return res
}

func Mountain(prev int, next int) int {
	res := 0
	if next == 0 && prev > 0 {
		res += 1
	}
	return res
}
