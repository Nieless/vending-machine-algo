package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"vending-machine/system"
)

const managerPassword = "222"

func main() {

	// initialize machine as manager
	err := initVendingMachine()
	if err != nil {
		log.Fatal(err)
	}

	inv := system.GetInvInstance()
	if inv == nil{
		log.Fatal("err: stock empty")
	}

	log.Println("Current inventory state of machine:")

	// current inventory state of vending machine
	for _, inv := range *inv {
		fmt.Println(inv.String())
	}

	vault := system.GetVaultInstance()
	if vault == nil{
		log.Fatal("err: vault empty")
	}

	log.Println()
	log.Println("Current vault state of machine:")

	// current vault state of vending machine
	for _, coin := range *vault {
		fmt.Println(coin.String())
	}

	log.Println()
	log.Println("vending machine is ready to use for customer")

	commands := system.GetCommands()
	log.Println("below are the available actions to perform with vending machine")
	for _, cmd := range commands {
		log.Println(cmd.String())
	}

	log.Println("please enter valid command")
	txn := new(system.Transaction)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		input = strings.TrimSpace(input)
		txn.Inv = inv
		txn.Vault = vault
		err , msg := executeInput(input, txn)
		if err != nil{
			fmt.Println(err.Error())
			continue
		}
		log.Println(msg)
	}

	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
}

// executeInput executes the Input to the system current state
func executeInput(input string, txn *system.Transaction) (error, string) {
	rString := ""
	cmd, err := system.NewCmdFromStr(input)
	if err != nil {
		return fmt.Errorf("%s", err), rString
	}

	switch cmd.Number {
	case 1:
		err, rString = txn.InsertCoins(cmd)
	case 2:
		err, rString = txn.ChooseItems(cmd)
	case 3:
		err, rString = txn.GetItems(cmd)
	case 4:
		err, rString = txn.ReturnCoins()
	case 5:
		err, rString = txn.GetReturnCoins()
	}

	return err, rString
}

func initVendingMachine() error {
	log.Println("please enter your manager password to start vending machine")
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input := scanner.Text()
		input = strings.TrimSpace(input)
		if input != managerPassword {
			return fmt.Errorf("error: you are not authorized")
		}
		break
	}

	return nil
}
