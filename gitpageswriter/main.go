// main.go
// This example demonstrates usage of child command

package main

import (
	"fmt"
	"os"

	"github.com/mkideal/cli"
)

func main() {
	if err := cli.Root(root,
		cli.Tree(help),
		cli.Tree(child),
		cli.Tree(child2),
	).Run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

var help = cli.HelpCommand("display help information")

// root command
type rootT struct {
	cli.Helper
	Name string `cli:"name" usage:"your name"`
}

var root = &cli.Command{
	Desc: "this is root command",
	// Argv is a factory function of argument object
	// ctx.Argv() is if Command.Argv == nil or Command.Argv() is nil
	Argv: func() interface{} { return new(rootT) },
	Fn: func(ctx *cli.Context) error {
		argv := ctx.Argv().(*rootT)
		ctx.String("Hello, root command, I am %s\n", argv.Name)
		return nil
	},
}

// child command
type childT struct {
	cli.Helper
	Name string `cli:"name" usage:"your name"`
}

type childT2 struct {
	cli.Helper
	Name string `cli:"name" usage:"your name"`
}

var child = &cli.Command{
	Name: "post",
	Desc: "this is a child command",
	Argv: func() interface{} { return new(childT) },
	Fn: func(ctx *cli.Context) error {
		argv := ctx.Argv().(*childT)
		ctx.String("Hello, child command, I am %s\n", argv.Name)
		return nil
	},
}

var child2 = &cli.Command{
	Name: "hama",
	Desc: "this is a hama command",
	Argv: func() interface{} { return new(childT2) },
	Fn: func(ctx *cli.Context) error {
		argv := ctx.Argv().(*childT2)
		ctx.String("Hello, hama command, I am %s\n", argv.Name)
		return nil
	},
}
