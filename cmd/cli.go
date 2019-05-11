// +build cli

// main application for the cli tool
package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/akosgarai/logicgates/cmd/cli"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	cliTool := cli.New()
	for {
		cliTool.Prompt()
		cmdString, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		err = cliTool.RunCommand(cmdString)
		if err != nil {
			cliTool.ErrorLog(err.Error())
		}
	}
}
