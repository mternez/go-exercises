package main

import (
	"fmt"
	"task/cli"
)

func main() {

	addCommand := cli.NewCommand(
		"add",
		func(arguments map[string]*cli.ArgumentValue) error {
			for key, value := range arguments {
				fmt.Print(key + " - > ")
				fmt.Println(value.Value)
			}
			return nil
		},
		cli.WithArgument(
			cli.WithName("title"),
			cli.WithHelper("Title of the task"),
			cli.WithOptional(false),
		),
		cli.WithArgument(
			cli.WithName("description"),
			cli.WithHelper("Description of the task"),
			cli.WithOptional(true),
		),
	)

	removeCommand := cli.NewCommand(
		"remove",
		func(arguments map[string]*cli.ArgumentValue) error {
			for key, value := range arguments {
				fmt.Print(key + " - > ")
				fmt.Println(value.Value)
			}
			return nil
		},
		cli.WithArgument(
			cli.WithName("title"),
			cli.WithHelper("Title of the task"),
			cli.WithOptional(false),
		),
	)

	runner := cli.NewCommandRunner(addCommand, removeCommand)

	errs := runner.Init()

	if len(errs) != 0 {
		for _, err := range errs {
			fmt.Print("ERR: ")
			fmt.Println(err)
		}
		return
	}

	errs = runner.Run()

	if len(errs) != 0 {
		for _, err := range errs {
			fmt.Print("ERR: ")
			fmt.Println(err)
		}
		return
	}
}
