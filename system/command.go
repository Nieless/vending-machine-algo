package system

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	commandInsertCoins    = "Insert Coins"
	commandChooseItems    = "Choose Items"
	commandGetItems       = "Get Items"
	commandReturnCoins    = "Return Coins"
	commandGetReturnCoins = "Get Return Coins"

	invalidInputErr = "error: cannot recognize input"
)

type Command struct {
	Name           string
	Number         int
	ArgumentNeeded bool
	Arguments      []int
}

var commands = map[int]Command{
	1: Command{
		Name:           commandInsertCoins,
		Number:         1,
		ArgumentNeeded: true,
	},
	2: Command{
		Name:           commandChooseItems,
		Number:         2,
		ArgumentNeeded: true,
	},
	3: Command{
		Name:           commandGetItems,
		Number:         3,
		ArgumentNeeded: false,
	},
	4: Command{
		Name:           commandReturnCoins,
		Number:         4,
		ArgumentNeeded: false,
	},
	5: Command{
		Name:           commandGetReturnCoins,
		Number:         5,
		ArgumentNeeded: false,
	},
}

// representation method
func (cmd *Command) String() string {
	cmdStr := fmt.Sprintf(`
		Command name : %s,
		Command number : %d`,
		cmd.Name, cmd.Number)

	return cmdStr
}

// NewCmdFromStr will generate a Command from user input string
func NewCmdFromStr(input string) (*Command, error) {
	inputSplitArr := strings.Fields(input) // input = "1 20"

	if len(inputSplitArr) > 2 {
		return nil, fmt.Errorf(invalidInputErr)
	}

	cmdInstruction := inputSplitArr[0]
	cmdNumber, err := strconv.Atoi(cmdInstruction)
	if err != nil {
		return nil, fmt.Errorf(invalidInputErr)
	}

	// if its a valid command
	if _, ok := commands[cmdNumber]; !ok {
		return nil, fmt.Errorf(invalidInputErr)
	}

	cmd := commands[cmdNumber]
	if cmd.ArgumentNeeded {
		argument, err := strconv.Atoi(inputSplitArr[1])
		if err != nil {
			return nil, fmt.Errorf("argument needed")
		}
		cmd.Arguments = append(cmd.Arguments, argument)
	}

	return &cmd, nil
}

// get all commands
func GetCommands() map[int]Command {
	return commands
}
