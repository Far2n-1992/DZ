package storage

import (
	"fmt"
	"os"
)

func WriteFiles(content []byte, name string) {
	file, err := os.Create(name)
	if err != nil {
		fmt.Println(err)
	}
	_, err = file.Write(content)
	defer file.Close()
	if err != nil {

		fmt.Println(err)
		return
	}
	fmt.Println("Запись успешна")
}
func ReadFile(name string) ([]byte, error) {
	data, err := os.ReadFile(name)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return data, nil
}
