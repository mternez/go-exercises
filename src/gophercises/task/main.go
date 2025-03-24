package main

import (
	"fmt"
	"task/cli"
	"task/manager"
)

func Print(errs []error) {
	for _, err := range errs {
		fmt.Printf("ERR: %s\n", err)
	}
}

func main() {

	mgr := manager.NewTaskManager("tasks.db")

	add := cli.NewCommand(
		"add",
		"Add a new task to your TODO list",
		func(arguments map[string]*cli.ArgumentValue) error {
			title := arguments["title"].Value
			task := &manager.Task{Title: title}
			description, hasDescription := arguments["description"]
			if hasDescription {
				task.Description = description.Value
			}
			err := mgr.Save(task)
			if err != nil {
				return err
			}
			fmt.Printf("\n'%s' added.\n", task.Title)
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

	remove := cli.NewCommand(
		"remove",
		"Removes a task from your TODO list",
		func(arguments map[string]*cli.ArgumentValue) error {
			title := arguments["title"].Value
			err := mgr.Remove(title)
			if err != nil {
				return err
			}
			fmt.Printf("\n'%s' has been removed.\n", title)
			return nil
		},
		cli.WithArgument(
			cli.WithName("title"),
			cli.WithHelper("Title of the task"),
			cli.WithOptional(false),
		),
	)

	do := cli.NewCommand(
		"do",
		"Mark a task on your TODO list as complete",
		func(arguments map[string]*cli.ArgumentValue) error {
			title := arguments["title"].Value
			task, err := mgr.FindByTitle(title)
			if err != nil {
				return err
			}
			task.Done = true
			err = mgr.Save(task)
			if err != nil {
				return err
			}
			fmt.Printf("\n'%s' has been marked done.\n", task.Title)
			return nil
		},
		cli.WithArgument(
			cli.WithName("title"),
			cli.WithHelper("Title of the task"),
			cli.WithOptional(false),
		),
	)

	list := cli.NewCommand(
		"list",
		"List all of your incomplete tasks",
		func(arguments map[string]*cli.ArgumentValue) error {
			tasks, err := mgr.FindToDo()
			if err != nil {
				return err
			}
			fmt.Printf("Tasks : \n\n")
			for _, task := range tasks {
				fmt.Printf("\t'%s'\t%s\n", task.Title, task.Description)
			}
			fmt.Printf("\n")
			return nil
		},
	)

	completed := cli.NewCommand(
		"completed",
		"List all of your completed tasks",
		func(arguments map[string]*cli.ArgumentValue) error {
			tasks, err := mgr.FindDone()
			if err != nil {
				return err
			}
			fmt.Printf("Tasks : \n\n")
			for _, task := range tasks {
				fmt.Printf("\t'%s'\t%s\n", task.Title, task.Description)
			}
			fmt.Printf("\n")
			return nil
		},
	)

	runner := cli.NewCommandRunner(
		"task",
		"task is a CLI for managing your TODOs.",
		add,
		remove,
		do,
		list,
		completed,
	)

	runner.Init()
	runner.Run()
	Print(runner.Errs)
}
