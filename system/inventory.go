package system

import (
	"fmt"
	"sync"
)

const (
	invalidItemErr = "error: invalid item selection"
)

type Item struct {
	SkuNumber    int
	Name         string
	Price        int
	AvailableQty int
}

type Inventory map[int]Item

var invInstance *Inventory
var once sync.Once

func GetInvInstance() *Inventory {
	once.Do(func() {
		inv := map[int]Item{
			1: {
				SkuNumber:    1,
				Name:         "Canned coffee",
				Price:        100,
				AvailableQty: 10,
			},
			2: {
				SkuNumber:    2,
				Name:         "Water PET bottle",
				Price:        150,
				AvailableQty: 10,
			},
			3: {
				SkuNumber:    3,
				Name:         "Sport drinks",
				Price:        200,
				AvailableQty: 10,
			},
		}
		invInstance = (*Inventory)(&inv)
	})

	return invInstance
}

// representation method
func (item *Item) String() string {
	itemStr := fmt.Sprintf(`
		Item name : %s,
		Item price : %d`,
		item.Name, item.Price)

	return itemStr
}

