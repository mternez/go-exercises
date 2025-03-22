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

func (c *command) printHelp() {
	fmt.Printf("%s : %s\n\n", c.name, c.description)
	fmt.Printf("Usage: %s ", c.name)
	for _, arg := range c.arguments {
		fmt.Printf("%s [%s] ", arg.name, arg.name)
	}
	fmt.Printf("\n\nArguments: \n")
	for _, arg := range c.arguments {
		if arg.optional {
			fmt.Printf("\t%s (Optional)\t%s\n", arg.name, arg.helper)
		} else {
			fmt.Printf("\t%s\t%s\n", arg.name, arg.helper)
		}
	}
}

type CommandRunner struct {
	name        string
	description string
	commands    map[string]*command
	selected    *command
	Errs        []error
	help        bool
}

func (runner *CommandRunner) PrintHelp() {
	fmt.Printf("%s\n\n", runner.description)
	fmt.Println("Usage:")
	fmt.Printf("\t%s [command]\n\n", runner.name)
	fmt.Println("Available commands:")
	for _, command := range runner.commands {
		fmt.Printf("\t%s\t%s\n", command.name, command.description)
	}
	fmt.Printf("\nUse \"%s [command] --help\" for more information about a command.\n", runner.name)
}

func NewCommandRunner(name string, description string, commands ...*command) *CommandRunner {

	runner := &CommandRunner{name: name, description: description, commands: make(map[string]*command), Errs: make([]error, 0)}

	for _, command := range commands {
		runner.commands[command.name] = command
	}

	return runner
}

func (runner *CommandRunner) Init() {

	if len(os.Args) < 2 {
		runner.help = true
		return
	}

	args := os.Args[1:]

	commandName := args[0]

	selectedCommand, ok := runner.commands[commandName]

	if !ok {
		runner.Errs = append(runner.Errs, errors.New("Command '"+commandName+"' does not exist."))
		return
	}

	runner.selected = selectedCommand

	args = args[1:]

	var currentArgument *argumentDescription
	for _, val := range args {

		if val == "--help" {
			runner.help = true
			return
		}

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
	validationErrs := runner.selected.validate()

	if len(validationErrs) > 0 {
		runner.Errs = append(runner.Errs, validationErrs...)
	}
}

func (runner *CommandRunner) Run() {

	if runner.help {
		if runner.selected != nil {
			runner.selected.printHelp()
			return
		}
		runner.PrintHelp()
		return
	}

	// Run the command
	if len(runner.Errs) == 0 {
		runErrors := runner.selected.run(runner.selected.values)
		if runErrors != nil {
			runner.Errs = append(runner.Errs, runErrors)
		}
	}
}
