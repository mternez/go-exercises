package cli

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
)

type CommandArgument struct {
	name      string
	valueType reflect.Kind
	nullable  bool
}

func (c *CommandArgument) Validate(value string) (bool, error) {
	var err error
	var valueValid bool = true
	nameValid := c.name != ""
	if !c.nullable {
		valueValid = value != ""
		if !valueValid {
			err = errors.New("Argument '" + c.name + "' has nil value '" + value + "' but is not declared nullable.")
		}
	}
	switch c.valueType {
	case reflect.Int:
		_, err = strconv.ParseInt(value, 10, 64)
		valueValid = err != nil
	case reflect.Bool:
		_, err = strconv.ParseBool(value)
		valueValid = err != nil
	default:
		valueValid = true
	}
	if !valueValid {
		err = errors.New("Argument '" + c.name + "' is declared as '" + c.valueType.String() + "'.")
	}
	return nameValid && valueValid, err
}

func NewCommandArgument(name string, valueType reflect.Kind, nullable bool) *CommandArgument {
	return &CommandArgument{name: name, valueType: valueType, nullable: nullable}
}

type Command struct {
	name      string
	arguments map[string]*CommandArgument
}

func NewCommand(name string) *Command {
	return &Command{name: name, arguments: make(map[string]*CommandArgument)}
}

func (c *Command) AddArgument(arg *CommandArgument) {
	c.arguments[arg.name] = arg
}

type ShellContext struct {
	// Context : Args should be a map (argname, argvalue)
	args []string
}

type Shell struct {
	context  *ShellContext
	commands map[string]*Command
}

func InitShell(commands []*Command) *Shell {

	var ctx *ShellContext

	// TODO : Constructing arguments should be more complexe
	// Exemple :
	// $ task add clean dishes
	// Added "clean dishes" to your task list.
	// Here it can't just be {"add", "clean"}
	// Though in this case it can be considered that we'll only ever have one argument
	if len(os.Args) > 1 {
		ctx = &ShellContext{args: os.Args[1:]}
	} else {
		ctx = &ShellContext{args: make([]string, 0)}
	}

	commandsMap := make(map[string]*Command)
	for _, command := range commands {
		commandsMap[command.name] = command
	}

	return &Shell{context: ctx, commands: commandsMap}
}

func (s *Shell) Run() error {
	commandName := s.context.args[0]
	command, ok := s.commands[commandName]
	if !ok {
		return errors.New("Command '" + commandName + "' does not exist.")
	}
	ind := 1
	for _, value := range command.arguments {
		_, err := value.Validate(s.context.args[ind])
		if err != nil {
			return err
		}
	}
	fmt.Println(command.name)
	return nil
}
