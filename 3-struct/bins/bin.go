package bins

import (
	"time"
)

type Bin struct {
	Id       string    `json:"id"`
	Private  bool      `json:"private"`
	CreateAt time.Time `json:"createAt"`
	Name     string    `json:"name"`
}

func CreateBin(id string, private bool, name string) *Bin {
	newBin := &Bin{
		Id:       id,
		Private:  private,
		CreateAt: time.Now(),
		Name:     name,
	}
	return newBin
}
