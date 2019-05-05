// +build cli

// main application for the cli tool
package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/fatih/color"
)

func welcome() {
	d := color.New(color.FgCyan, color.Bold)
	d.Println("LogicGate cli-tool")
}

func printClolorsForTesting() {
	color.Blue("LogicGate cli-tool")
	color.Cyan("LogicGate cli-tool")
	color.Magenta("LogicGate cli-tool")
	color.White("LogicGate cli-tool")
	color.Yellow("LogicGate cli-tool")
	color.Green("LogicGate cli-tool")
	color.Red("LogicGate cli-tool")
}

func prompt() {
	t := time.Now()
	d := color.New(color.FgRed, color.Bold, color.BgBlue)
	d.Printf("[ %s ] cli-tool $ ", t.Format("2006-01-02 15:04:05"))
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	welcome()
	printClolorsForTesting()
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
	arrCommandStr := strings.Fields(commandStr)
	switch arrCommandStr[0] {
	case "exit":
		os.Exit(0)
	case "help":
		return help()
	}
	cmd := exec.Command("echo", commandStr)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
func help() error {
	fmt.Println("Commands:")
	fmt.Println(" - help")
	fmt.Println(" - exit")
	return nil
}
