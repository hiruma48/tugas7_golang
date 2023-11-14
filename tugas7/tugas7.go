package main

import (
	"fmt"
	"reflect"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)

	var nama = "ari"
	var umur = 20

	go tampilkan_nama(5, nama)
	tampilkan_nama(umur, nama)

	// var input string
	// fmt.Scanln(&input)

	// go tampilkan_umur(2,umur)
	// tampilkan_umur(2,umur)

	// var input int
	// fmt.Scanln(&input)

	var reflectnama = reflect.ValueOf(nama)
	var reflectumur = reflect.ValueOf(umur)

	fmt.Println("tipe variabele :", reflectnama.Type())
	if reflectnama.Kind() == reflect.String {
		fmt.Println("nama :", reflectnama.String())
	}

	fmt.Println("tipe variabele :", reflectumur.Type())
	if reflectumur.Kind() == reflect.Int {
		fmt.Println("umur :", reflectumur.Int())
	}

}

func tampilkan_nama(umur int, nama string) {
	for i := 0; i < umur; i++ {
		fmt.Println((i + 1), nama)
		fmt.Println((i + 1), umur)
	}
}
