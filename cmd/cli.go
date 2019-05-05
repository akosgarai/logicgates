// +build cli

// main application for the cli tool
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
)

func welcome() {
	d := color.New(color.FgCyan, color.Bold)
	d.Println("LogicGate cli-tool")
}

func printClolorsForTesting() error {
	color.Blue("LogicGate cli-tool")
	color.Cyan("LogicGate cli-tool")
	color.Magenta("LogicGate cli-tool")
	color.White("LogicGate cli-tool")
	color.Yellow("LogicGate cli-tool")
	color.Green("LogicGate cli-tool")
	color.Red("LogicGate cli-tool")
	return nil
}

func prompt() {
	t := time.Now()
	d := color.New(color.FgRed, color.Bold, color.BgBlue)
	d.Printf("[ %s ] cli-tool $ ", t.Format("2006-01-02 15:04:05"))
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	welcome()
	help()
	for {
		prompt()
		cmdString, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		err = runCommand(cmdString)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}
func runCommand(commandStr string) error {
	commandStr = strings.TrimSuffix(commandStr, "\n")
	if commandStr == "" {
		return nil
	}
	arrCommandStr := strings.Fields(commandStr)
	switch arrCommandStr[0] {
	case "exit":
		os.Exit(0)
	case "colors":
		return printClolorsForTesting()
	case "help":
		return help()
	}
	return invalidCommand(commandStr)
}
func help() error {
	d := color.New(color.FgBlack, color.Bold, color.BgYellow)
	d.Println("Commands:")
	d.Println(" - help")
	d.Println(" - exit")
	d.Println(" - colors")
	fmt.Println("")
	return nil
}
func invalidCommand(cmdString string) error {
	d := color.New(color.FgBlack, color.Bold, color.BgRed)
	d.Println("Invalid command! The following command does not exists.")
	d.Println(cmdString)
	d.Println("Type help for the helpmenu.")
	fmt.Println("")
	return nil
}
