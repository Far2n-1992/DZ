package main

import (
	"fmt"

	"main.go/bins"
	"main.go/file"
)

func main() {
	binlist, err := bins.NewBinlist()
	if err != nil {
		fmt.Println(err)
	}
	createBin(binlist)
	data, err := file.ReadFile("file.txt") // проверка работоспособности
	if err != nil {                        // проверка работоспособности
		fmt.Println(err) // проверка работоспособности
	}
	fmt.Println(string(data)) // проверка работоспособности
}
func createBin(binlist *bins.BinList) {
	var a string // переменные для входящих данных
	var b bool   // переменные для входящих данных
	var c string // переменные для входящих данных
	a = "test"
	b = true
	c = "test2"
	bin := bins.CreateBin(a, b, c)
	binlist.AddNewBin(*bin)
}
