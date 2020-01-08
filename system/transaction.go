package system

import "fmt"

const (
	CoinInserted    = 10
	ItemChoose      = 20
	GetItems        = 30
	CoinReturned    = 40
	GetReturnedCoin = 50
)

type Transaction struct {
	Items         map[int]int
	InsertedCoins map[int]int
	ReturnedCoins map[int]int
	State         int

	Inv   *Inventory
	Vault *Vault
}

// update vault
func (txn *Transaction) InsertCoins(cmd *Command) (error, string) {
	if txn.State >= GetItems {
		return fmt.Errorf("you can not run this command now as transaction is in progress, please collect item or coin before inseting again"), ""
	}

	for _, coin := range cmd.Arguments {

		if _, ok := txn.InsertedCoins[coin]; !ok {
			txn.InsertedCoins = map[int]int{coin: 1,}
		} else {
			txn.InsertedCoins[coin] = txn.InsertedCoins[coin] + 1
		}

		if txn.Vault != nil {
			vault := *txn.Vault
			v := vault[coin]
			fmt.Println(v, "vault")
			v.AvailableQty = v.AvailableQty + 1
		}
	}
	txn.State = CoinInserted

	cStr := "Inserted coins: "
	for coinType, qty := range txn.InsertedCoins {
		coinStr := fmt.Sprintf(`
			Coin type : %d
			Qty : %d`,
			coinType, qty,
		)
		cStr = cStr + " " + coinStr
	}

	return nil, cStr
}

// update inventory
func (txn *Transaction) ChooseItems(cmd *Command) (error, string) {
	fmt.Println(txn, "ddddd")
	if txn.State != CoinInserted {
		return fmt.Errorf("please run command 1 to insert coins before choosing items"), ""
	}

	for _, itemSkuNumber := range cmd.Arguments {

		if _, ok := txn.Items[itemSkuNumber]; !ok {
			txn.Items = map[int]int{itemSkuNumber: 1,}
		} else {
			txn.Items[itemSkuNumber] = txn.Items[itemSkuNumber] + 1
		}

		if txn.Inv != nil {
			inv := *txn.Inv
			item := inv[itemSkuNumber]
			item.AvailableQty = item.AvailableQty - 1
		}
	}

	iStr := "Items Selected:"
	for itemNumber, qty := range txn.Items {
		inv := *txn.Inv
		item := inv[itemNumber]
		itemStr := fmt.Sprintf(`
			Item : %s
			Qty : %d`,
			item.Name, qty,
		)
		iStr = iStr + " " + itemStr
	}

	txn.State = ItemChoose
	return nil, iStr
}

func (txn *Transaction) GetItems(cmd *Command) (error, string) {
	if txn.State != ItemChoose {
		return fmt.Errorf("please run command 2 to choose items"), ""
	}

	iStr := "Items :"
	for itemNumber, qty := range txn.Items {
		inv := *txn.Inv
		item := inv[itemNumber]
		itemStr := fmt.Sprintf(`
			Item : %s
			Qty : %d`,
			item.Name, qty,
		)
		iStr = iStr + " " + itemStr
	}

	txn.State = GetItems
	return nil, iStr
}

func (txn *Transaction) ReturnCoins() (error, string) {
	for _, coin := range txn.ReturnedCoins {
		txn.ReturnedCoins[coin] = + 1
		if txn.Vault != nil {
			vault := *txn.Vault
			v := vault[coin]
			v.AvailableQty = v.AvailableQty - 1
		}
	}
	txn.State = CoinReturned
	return nil, ""
}

// asked for action
func (txn *Transaction) GetReturnCoins() (error, string) {
	if txn.State != CoinReturned {
		return fmt.Errorf("please run command 4 to get returned coins"), ""
	}

	cStr := "Returned Coins: "
	for coinType, qty := range txn.ReturnedCoins {
		coinStr := fmt.Sprintf(`
			Coin type : %d
			Qty : %d`,
			coinType, qty,
		)
		cStr = cStr + " " + coinStr
	}

	txn.State = GetReturnedCoin
	return nil, cStr
}
