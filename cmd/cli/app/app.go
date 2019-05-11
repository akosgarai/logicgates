package cli

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/akosgarai/logicgates/basic"
	"github.com/fatih/color"
)

type Command struct {
	Name        string
	Description string
}

type Cli struct {
	prompt   *color.Color
	welcome  *color.Color
	help     *color.Color
	invalid  *color.Color
	commands []Command
}

func New() *Cli {
	c := &Cli{
		prompt:  color.New(color.FgRed, color.Bold, color.BgBlue),
		welcome: color.New(color.FgCyan, color.Bold),
		help:    color.New(color.FgBlack, color.Bold, color.BgYellow),
		invalid: color.New(color.FgBlack, color.Bold, color.BgRed),
		commands: []Command{
			{"help", "Print out the help menu."},
			{"colors", "Print out a sample text for checking the colors that we can use."},
			{"exit", "Quit from the application."},
			{"and", "The and application demonstrates the And gate."},
			{"or", "The or application demonstrates the Or gate."},
			{"nor", "The nor application demonstrates the Nor gate."},
			{"not", "The not application demonstrates the Not gate."},
			{"nand", "The nand application demonstrates the Nand gate."},
			{"and3", "The and3 application demonstrates the And3 gate."},
			{"xor", "The xor application demonstrates the Xor gate."},
			{"xnor", "The xnor application demonstrates the Xnor gate."},
			{"halfadd", "The halfadd application demonstrates the HalfAdd gate."},
			{"fulladd", "The fulladd application demonstrates the FullAdd gate."},
		},
	}
	c.Welcome()
	c.Help()
	return c
}

func (c *Cli) Welcome() {
	c.welcome.Println("LogicGate cli-tool")
}
func (c *Cli) ErrorLog(msg string) {
	c.invalid.Println(msg)
	fmt.Println("")
}

func (c *Cli) printClolorsForTesting() error {
	color.Blue("LogicGate cli-tool")
	color.Cyan("LogicGate cli-tool")
	color.Magenta("LogicGate cli-tool")
	color.White("LogicGate cli-tool")
	color.Yellow("LogicGate cli-tool")
	color.Green("LogicGate cli-tool")
	color.Red("LogicGate cli-tool")
	c.prompt.Println("prompt style text")
	c.welcome.Println("welcome style text")
	c.help.Println("help style text")
	c.invalid.Println("invalid style text")
	fmt.Println("")
	return nil
}

func (c *Cli) Prompt() {
	c.prompt.Printf("[ %s ] cli-tool $ ", time.Now().Format("2006-01-02 15:04:05"))
}
func (c *Cli) Help() error {
	c.help.Println("Commands:")
	for _, item := range c.commands {
		c.help.Println(" - ", item.Name, " : ", item.Description)
	}
	fmt.Println("")
	return nil
}
func (c *Cli) and(cmdString string) error {
	commandStr := strings.TrimSuffix(cmdString, "\n")
	arrCommandStr := strings.Fields(commandStr)
	if len(arrCommandStr) != 3 {
		return errors.New("Wrong number of arguments. Try this: and 0 0")
	}
	arg1 := arrCommandStr[1] == "1"
	arg2 := arrCommandStr[2] == "1"
	c.welcome.Println("And(", arg1, ",", arg2, ") =", logicgates.And(arg1, arg2))
	fmt.Println("")
	return nil
}
func (c *Cli) or(cmdString string) error {
	commandStr := strings.TrimSuffix(cmdString, "\n")
	arrCommandStr := strings.Fields(commandStr)
	if len(arrCommandStr) != 3 {
		c.invalid.Println("Wrong number of arguments. Try this: or 0 0")
	}
	arg1 := arrCommandStr[1] == "1"
	arg2 := arrCommandStr[2] == "1"
	c.welcome.Println("Or(", arg1, ",", arg2, ") =", logicgates.Or(arg1, arg2))
	fmt.Println("")
	return nil
}
func (c *Cli) nor(cmdString string) error {
	commandStr := strings.TrimSuffix(cmdString, "\n")
	arrCommandStr := strings.Fields(commandStr)
	if len(arrCommandStr) != 3 {
		c.invalid.Println("Wrong number of arguments. Try this: nor 0 0")
	}
	arg1 := arrCommandStr[1] == "1"
	arg2 := arrCommandStr[2] == "1"
	c.welcome.Println("Nor(", arg1, ",", arg2, ") =", logicgates.Nor(arg1, arg2))
	fmt.Println("")
	return nil
}
func (c *Cli) not(cmdString string) error {
	commandStr := strings.TrimSuffix(cmdString, "\n")
	arrCommandStr := strings.Fields(commandStr)
	if len(arrCommandStr) != 2 {
		c.invalid.Println("Wrong number of arguments. Try this: not 0")
	}
	arg1 := arrCommandStr[1] == "1"
	c.welcome.Println("Not(", arg1, ") =", logicgates.Not(arg1))
	fmt.Println("")
	return nil
}
func (c *Cli) nand(cmdString string) error {
	commandStr := strings.TrimSuffix(cmdString, "\n")
	arrCommandStr := strings.Fields(commandStr)
	if len(arrCommandStr) != 3 {
		c.invalid.Println("Wrong number of arguments. Try this: nand 0 0")
	}
	arg1 := arrCommandStr[1] == "1"
	arg2 := arrCommandStr[2] == "1"
	c.welcome.Println("Nand(", arg1, ",", arg2, ") =", logicgates.Nand(arg1, arg2))
	fmt.Println("")
	return nil
}
func (c *Cli) and3(cmdString string) error {
	commandStr := strings.TrimSuffix(cmdString, "\n")
	arrCommandStr := strings.Fields(commandStr)
	if len(arrCommandStr) != 4 {
		c.invalid.Println("Wrong number of arguments. Try this: and3 0 0 0")
	}
	arg1 := arrCommandStr[1] == "1"
	arg2 := arrCommandStr[2] == "1"
	arg3 := arrCommandStr[3] == "1"
	c.welcome.Println("And3(", arg1, ",", arg2, ",", arg3, ") =", logicgates.And3(arg1, arg2, arg3))
	fmt.Println("")
	return nil
}
func (c *Cli) xor(cmdString string) error {
	commandStr := strings.TrimSuffix(cmdString, "\n")
	arrCommandStr := strings.Fields(commandStr)
	if len(arrCommandStr) != 3 {
		c.invalid.Println("Wrong number of arguments. Try this: xor 0 0")
	}
	arg1 := arrCommandStr[1] == "1"
	arg2 := arrCommandStr[2] == "1"
	c.welcome.Println("Xor(", arg1, ",", arg2, ") =", logicgates.Xor(arg1, arg2))
	fmt.Println("")
	return nil
}
func (c *Cli) xnor(cmdString string) error {
	commandStr := strings.TrimSuffix(cmdString, "\n")
	arrCommandStr := strings.Fields(commandStr)
	if len(arrCommandStr) != 3 {
		c.invalid.Println("Wrong number of arguments. Try this: xnor 0 0")
	}
	arg1 := arrCommandStr[1] == "1"
	arg2 := arrCommandStr[2] == "1"
	c.welcome.Println("Xnor(", arg1, ",", arg2, ") =", logicgates.Xnor(arg1, arg2))
	fmt.Println("")
	return nil
}
func (c *Cli) halfadd(cmdString string) error {
	commandStr := strings.TrimSuffix(cmdString, "\n")
	arrCommandStr := strings.Fields(commandStr)
	if len(arrCommandStr) != 3 {
		c.invalid.Println("Wrong number of arguments. Try this: halfadd 0 0 0")
	}
	arg1 := arrCommandStr[1] == "1"
	arg2 := arrCommandStr[2] == "1"
	o1, o2 := logicgates.HalfAdd(arg1, arg2)
	c.welcome.Println("HalfAdd(", arg1, ",", arg2, ") =", o1, ",", o2)
	fmt.Println("")
	return nil
}
func (c *Cli) fulladd(cmdString string) error {
	commandStr := strings.TrimSuffix(cmdString, "\n")
	arrCommandStr := strings.Fields(commandStr)
	if len(arrCommandStr) != 4 {
		c.invalid.Println("Wrong number of arguments. Try this: fulladd 0 0 0")
	}
	arg1 := arrCommandStr[1] == "1"
	arg2 := arrCommandStr[2] == "1"
	arg3 := arrCommandStr[3] == "1"
	o1, o2 := logicgates.FullAdd(arg1, arg2, arg3)
	c.welcome.Println("FullAdd(", arg1, ",", arg2, ",", arg3, ") =", o1, ",", o2)
	fmt.Println("")
	return nil
}
func (c *Cli) invalidCommand(cmdString string) error {
	c.invalid.Println("Invalid command! The following command does not exist.")
	c.invalid.Println(cmdString)
	c.invalid.Println("Type help for the helpmenu.")
	fmt.Println("")
	return nil
}
func (c Cli) RunCommand(commandStr string) error {
	commandStr = strings.TrimSuffix(commandStr, "\n")
	if commandStr == "" {
		return nil
	}
	arrCommandStr := strings.Fields(commandStr)
	switch arrCommandStr[0] {
	case "exit":
		os.Exit(0)
	case "colors":
		return c.printClolorsForTesting()
	case "help":
		return c.Help()
	case "and":
		return c.and(commandStr)
	case "or":
		return c.or(commandStr)
	case "nor":
		return c.nor(commandStr)
	case "not":
		return c.not(commandStr)
	case "nand":
		return c.nand(commandStr)
	case "and3":
		return c.and3(commandStr)
	case "xor":
		return c.xor(commandStr)
	case "xnor":
		return c.xnor(commandStr)
	case "halfadd":
		return c.halfadd(commandStr)
	case "fulladd":
		return c.fulladd(commandStr)
	}
	return c.invalidCommand(commandStr)
}
