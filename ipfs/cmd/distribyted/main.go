package main

import (
	"fmt"
	"os"

	"github.com/ajnavarro/distribyted/cmd/distribyted/command"

	"github.com/jessevdk/go-flags"
	"github.com/sirupsen/logrus"
)

const (
	name = "distribyted"
)

func main() {
	parser := flags.NewNamedParser(name, flags.Default)

	parser.UnknownOptionHandler = func(option string, arg flags.SplitArgument, args []string) ([]string, error) {
		if option != "g" {
			return nil, fmt.Errorf("unknown flag `%s'", option)
		}

		if len(args) == 0 {
			return nil, fmt.Errorf("unknown flag `%s'", option)
		}

		return append(append(args, "-d"), args[0]), nil
	}

	_, err := parser.AddCommand("server",
		command.ServerDescription,
		command.ServerHelp,
		&command.Server{},
	)
	if err != nil {
		logrus.Fatal(err)
	}

	_, err = parser.Parse()
	if err != nil {
		if e, ok := err.(*flags.Error); ok && e.Type == flags.ErrCommandRequired {
			parser.WriteHelp(os.Stdout)
		}

		os.Exit(1)
	}
}
