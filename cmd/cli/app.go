package cli

import (
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
		},
	}
	c.Welcome()
	c.Help()
	return c
}

func (c *Cli) Welcome() {
	c.welcome.Println("LogicGate cli-tool")
}

func (c *Cli) PrintClolorsForTesting() error {
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
func (c *Cli) And(cmdString string) error {
	commandStr := strings.TrimSuffix(cmdString, "\n")
	arrCommandStr := strings.Fields(commandStr)
	if len(arrCommandStr) != 3 {
		c.invalid.Println("Wrong number of arguments. Try this: add 0 0")
	}
	arg1 := arrCommandStr[1] == "1"
	arg2 := arrCommandStr[2] == "1"
	c.welcome.Println("And(", arg1, ",", arg2, ") =", logicgates.And(arg1, arg2))
	fmt.Println("")
	return nil
}
func (c *Cli) InvalidCommand(cmdString string) error {
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
		return c.PrintClolorsForTesting()
	case "help":
		return c.Help()
	case "and":
		return c.And(commandStr)
	}
	return c.InvalidCommand(commandStr)
}
