package cli

import (
	"errors"
	"fmt"
	"os"
)

type argumentDescriptionOption interface {
	apply(command *argumentDescription)
}

type argumentNameOption struct {
	name string
}

func (opt *argumentNameOption) apply(command *argumentDescription) {
	command.name = opt.name
}

type argumentHelperOption struct {
	helper string
}

func (opt *argumentHelperOption) apply(command *argumentDescription) {
	command.helper = opt.helper
}

type argumentOptionalOption struct {
	optional bool
}

func (opt *argumentOptionalOption) apply(command *argumentDescription) {
	command.optional = opt.optional
}

type argumentNullableOption struct {
	nullable bool
}

func (opt *argumentNullableOption) apply(command *argumentDescription) {
	command.nullable = opt.nullable
}

type argumentValidatorOption struct {
	validator func(v string) error
}

func (opt *argumentValidatorOption) apply(command *argumentDescription) {
	command.validator = opt.validator
}

func WithName(name string) argumentDescriptionOption {
	return &argumentNameOption{name: name}
}

func WithHelper(helper string) argumentDescriptionOption {
	return &argumentHelperOption{helper: helper}
}

func WithOptional(optional bool) argumentDescriptionOption {
	return &argumentOptionalOption{optional: optional}
}

func WithValidator(validator func(v string) error) argumentDescriptionOption {
	return &argumentValidatorOption{validator: validator}
}

func WithNullable(nullable bool) argumentDescriptionOption {
	return &argumentNullableOption{nullable: nullable}
}

type argumentDescription struct {
	name      string
	helper    string
	optional  bool
	nullable  bool
	validator func(v string) error
}

type ArgumentValue struct {
	Name  string
	Value string
}

type command struct {
	name        string
	description string
	arguments   map[string]*argumentDescription
	values      map[string]*ArgumentValue
	run         func(arguments map[string]*ArgumentValue) error
}

func WithArgument(opts ...argumentDescriptionOption) *argumentDescription {
	argumentDescription := &argumentDescription{}
	for _, opt := range opts {
		opt.apply(argumentDescription)
	}
	return argumentDescription
}

func NewCommand(name string, description string, run func(arguments map[string]*ArgumentValue) error, arguments ...*argumentDescription) *command {

	command := &command{name: name, description: description, run: run, arguments: make(map[string]*argumentDescription), values: make(map[string]*ArgumentValue)}
	for _, arg := range arguments {
		command.arguments[arg.name] = arg
	}
	return command
}

func (description *argumentDescription) validate(value string) []error {

	errs := make([]error, 0)

	if description.name == "" {
		errs = append(errs, errors.New("'"+description.name+"' is not a valid command argument name."))
	}

	if !description.nullable && value == "" {
		errs = append(errs, errors.New("'"+description.name+"' must have a value."))
	}

	if description.validator != nil {
		errs = append(errs, description.validator(value))
	}

	return errs
}

func (c *command) validate() []error {
	errs := make([]error, 0)
	for _, arg := range c.arguments {
		argValue := c.values[arg.name]
		if argValue == nil {
			if !arg.optional {
				errs = append(errs, errors.New("'"+arg.name+"' must be set."))
			}
		} else {
			errs = append(errs, arg.validate(argValue.Value)...)
		}
	}
	return errs
}

type CommandRunner struct {
	name        string
	description string
	commands    map[string]*command
	selected    *command
}

func (runner *CommandRunner) PrintHelp() {
	fmt.Println(runner.description)
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Printf("\t%s [command]\n", runner.name)
	fmt.Println("Available commands:")
	for _, command := range runner.commands {
		fmt.Printf("\t%s\t%s\n", command.name, command.description)
	}
	fmt.Printf("Use \"%s [command] --help\" for more information about a command.\n", runner.name)
}

func NewCommandRunner(name string, description string, commands ...*command) *CommandRunner {

	runner := &CommandRunner{name: name, description: description, commands: make(map[string]*command)}

	for _, command := range commands {
		runner.commands[command.name] = command
	}

	return runner
}

func (runner *CommandRunner) Init() []error {

	errs := make([]error, 0)

	if len(os.Args) < 2 {
		return append(errs, errors.New("No command provided."))
	}

	args := os.Args[1:]

	commandName := args[0]

	selectedCommand, ok := runner.commands[commandName]

	if !ok {
		return append(errs, errors.New("Command '"+selectedCommand.name+"' does not exist."))
	}

	runner.selected = selectedCommand

	args = args[1:]

	var currentArgument *argumentDescription
	for _, val := range args {
		argument := runner.selected.arguments[val]
		if argument != nil {
			currentArgument = argument
			runner.selected.values[currentArgument.name] = &ArgumentValue{Name: currentArgument.name}
		} else if currentArgument != nil {
			currentValue := runner.selected.values[currentArgument.name]
			currentValue.Value = currentValue.Value + val
		}
	}

	// Validate each command argument
	errs = runner.selected.validate()

	return errs
}

func (runner *CommandRunner) Run() []error {

	errs := make([]error, 0)

	// Run the command
	err := runner.selected.run(runner.selected.values)

	if err != nil {
		errs = append(errs, err)
	}

	return errs
}
