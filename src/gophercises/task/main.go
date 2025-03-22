package main

import (
	"fmt"
	"task/cli"
)

func Print(errs []error) {
	for _, err := range errs {
		fmt.Print("ERR: ")
		fmt.Println(err)
	}
}

func main() {

	add := cli.NewCommand(
		"add",
		"Add a new task to your TODO list",
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

	do := cli.NewCommand(
		"do",
		"Mark a task on your TODO list as complete",
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

	runner := cli.NewCommandRunner(
		"task",
		"task is a CLI for managing your TODOs.",
		add,
		do,
	)

	runner.Init()
	runner.Run()
	Print(runner.Errs)
}
