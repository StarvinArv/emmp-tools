package main

import (
	"fmt"
	"os"
)

// Command represents a CLI command.
type Command struct {
	Use   string
	Short string
	RunFn func(args []string) error
	subs  []*Command
}

// AddCommand registers a subcommand.
func (c *Command) AddCommand(sub *Command) {
	c.subs = append(c.subs, sub)
}

// Execute runs the command.
func (c *Command) Execute() error {
	args := os.Args[1:]
	return c.run(args)
}

func (c *Command) run(args []string) error {
	if len(args) > 0 {
		for _, sub := range c.subs {
			if sub.Use == args[0] {
				return sub.run(args[1:])
			}
		}
	}

	if c.RunFn != nil {
		return c.RunFn(args)
	}

	// Print help.
	fmt.Printf("Usage: %s <command>\n\n", c.Use)
	fmt.Println("Commands:")
	for _, sub := range c.subs {
		fmt.Printf("  %-15s %s\n", sub.Use, sub.Short)
	}
	return nil
}

// Global flags.
var (
	flagServer string
	flagToken  string
	flagOutput string
)

func rootCmd() *Command {
	root := &Command{
		Use:   "emmp",
		Short: "Emmp developer CLI",
	}

	// Version command.
	root.AddCommand(&Command{
		Use:   "version",
		Short: "Print version",
		RunFn: func(_ []string) error {
			fmt.Printf("emmp version %s\n", version)
			return nil
		},
	})

	// Policy subcommand group.
	policyCmd := &Command{
		Use:   "policy",
		Short: "Cedar policy tooling commands",
	}
	policyCmd.AddCommand(validateCmd())
	policyCmd.AddCommand(testCmd())
	policyCmd.AddCommand(lintCmd())
	policyCmd.AddCommand(scanCmd())
	policyCmd.AddCommand(analyzeCmd())

	root.AddCommand(policyCmd)

	return root
}
