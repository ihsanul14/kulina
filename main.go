package main

import (
	"fmt"
)

func main() {
	var soal int
	fmt.Print("Masukkan nomor soal : ")
	fmt.Scanln(&soal)

	if soal == 1 {
		fmt.Println(Socks())
	} else if soal == 2 {
		fmt.Println(Step())
	} else if soal == 3 {
		DetailsNumber()
	} else if soal == 4 {
		fmt.Println(CountLamp())
	} else {
		fmt.Println("Invalid input soal")
	}

}
