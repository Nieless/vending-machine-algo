package system

import (
	"fmt"
	"sync"
)

type Coin struct {
	Number  int
	Type    int
	AvailableQty int
}

type Vault map[int]Coin

var vaultInstance *Vault
var vaultOnce sync.Once

func GetVaultInstance() *Vault {
	vaultOnce.Do(func() {
		vault := map[int]Coin{
			10: {
				Number:    1,
				Type:        10,
				AvailableQty: 10,
			},
			20: {
				Number:    2,
				Type:        20,
				AvailableQty: 10,
			},
			50: {
				Number:    3,
				Type:        50,
				AvailableQty: 5,
			},
			100: {
				Number:    4,
				Type:        100,
				AvailableQty: 5,
			},
			500: {
				Number:    5,
				Type:        500,
				AvailableQty: 2,
			},
		}
		vaultInstance = (*Vault)(&vault)
	})

	return vaultInstance
}

// representation method
func (c *Coin) String() string {
	cStr := fmt.Sprintf(`
		Coin Type : %d,
		Available Qty : %d`,
		c.Type, c.AvailableQty)

	return cStr
}
