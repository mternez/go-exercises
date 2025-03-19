package main

import (
	"fmt"
	"reflect"
	"task/cli"
)

func main() {
	commands := make([]*cli.Command, 1)

	commands[0] = cli.NewCommand("add")

	commandArgument := cli.NewCommandArgument("name", reflect.Int, false)
	commands[0].AddArgument(commandArgument)

	shell := cli.InitShell(commands)
	err := shell.Run()
	if err != nil {
		fmt.Println(err)
		return
	}
}
