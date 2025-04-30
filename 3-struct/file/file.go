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
	if IsJSONFile(name) {
		return data, nil
	} else {
		return nil, errors.New("Invalid_format_file")
	}
}
func IsJSONFile(name string) bool {
	return strings.HasSuffix(name, ".json")
}
