package main

import (
	"fmt"
	"task/cli"
)

func Run(command *cli.Command) (bool, error) {
	fmt.Println(command.Name)
	return false, nil
}

func main() {

	addCommand := cli.NewCommand(
		"add",
		func(arguments map[string]*cli.CommandArgument) (bool, error) {
			for key, value := range arguments {
				fmt.Print(key + " - > ")
				fmt.Println(value.Value)
			}
			return false, nil
		},
		cli.NewCommandArgument("name", false),
		cli.NewCommandArgument("description", true),
	)

	removeCommand := cli.NewCommand(
		"remove",
		func(arguments map[string]*cli.CommandArgument) (bool, error) {
			for key, value := range arguments {
				fmt.Print(key + " - > ")
				fmt.Println(value.Value)
			}
			return false, nil
		},
		cli.NewCommandArgument("name", false),
	)

	shell := cli.InitShell(addCommand, removeCommand)
	err := shell.Run()
	if err != nil {
		fmt.Println(err)
		return
	}
}
