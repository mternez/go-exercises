package cli

import (
	"errors"
	"os"
)

type CommandRunner interface {
	Run(command *Command) (bool, error)
}

type CommandArgument struct {
	Name     string
	nullable bool
	Value    string
}

type Command struct {
	Name      string
	arguments map[string]*CommandArgument
	Run       func(arguments map[string]*CommandArgument) error
}

func NewCommandArgument(name string, nullable bool) *CommandArgument {
	return &CommandArgument{Name: name, nullable: nullable}
}

func NewCommand(name string, runner func(arguments map[string]*CommandArgument) error, arguments ...*CommandArgument) *Command {

	command := &Command{Name: name, Run: runner, arguments: make(map[string]*CommandArgument)}
	for _, argument := range arguments {
		command.arguments[argument.Name] = argument
	}
	return command
}

func (c *CommandArgument) Validate() (bool, error) {
	var err error
	var valueValid bool = true
	nameValid := c.Name != ""
	if !c.nullable {
		valueValid = c.Value != ""
		if !valueValid {
			err = errors.New("'" + c.Name + "' must be set.")
		}
	}
	return nameValid && valueValid, err
}

func (c *Command) AddArgument(arg *CommandArgument) {
	c.arguments[arg.Name] = arg
}

type Shell struct {
	commands map[string]*Command
}

func InitShell(commands ...*Command) *Shell {

	commandsMap := make(map[string]*Command)
	for _, command := range commands {
		commandsMap[command.Name] = command
	}

	return &Shell{commands: commandsMap}
}

func (s *Shell) Run() error {

	args := os.Args[1:]

	commandName := args[0]

	command, ok := s.commands[commandName]

	if !ok {
		return errors.New("Command '" + commandName + "' does not exist.")
	}

	args = args[1:]

	var currentArgumentName string

	for _, val := range args {

		argument, isArgument := command.arguments[val]

		if isArgument {
			currentArgumentName = argument.Name
		} else {
			command.arguments[currentArgumentName].Value = command.arguments[currentArgumentName].Value + " " + val
		}
	}

	for _, command := range command.arguments {
		_, err := command.Validate()
		if err != nil {
			panic(err)
		}
	}

	return command.Run(command.arguments)
}
