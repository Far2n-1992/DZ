package bins

import (
	"encoding/json"
	"fmt"

	"main.go/storage"
)

type BinList struct {
	Bins []Bin `json:"bins"`
}

func (binlist *BinList) AddNewBin(bin Bin) {
	binlist.Bins = append(binlist.Bins, bin)
	binlist.DataSave()
}
func (binlist *BinList) ToBytes() ([]byte, error) {
	file, err := json.Marshal(binlist)
	return file, err
}
func (binlist *BinList) DataSave() {
	data, err := binlist.ToBytes()
	if err != nil {
		fmt.Println("Не удалось разобрать файл json")
		return
	}
	storage.WriteFiles(data, "binlist.json")

}
func NewBinlist() (*BinList, error) {
	file, err := storage.ReadFile("binlist.json")
	if err != nil {
		return &BinList{
			Bins: []Bin{},
		}, nil
	}
	var vault BinList
	err = json.Unmarshal(file, &vault)
	if err != nil {
		return nil, err
		// fmt.Println("Не удалось преоброзовать файл json")
	}
	return &vault, nil
}
