package file

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func ReadFile(name string) ([]byte, error) {
	data, err := os.ReadFile(name)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	a := strings.Contains(name, ".json") // проверка на расширение файла json
	if a {
		return data, nil
	} else {
		return nil, errors.New("Invalid_format_file")
	}

}
