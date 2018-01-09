package cmd

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/google/subcommands"
)

func Exec(version string) subcommands.ExitStatus {
	v := flag.Bool("v", false, "Print version.")
	flag.Parse()
	if *v {
		fmt.Printf("psb-rotatorc v%s\n", version)
		os.Exit(0)
	}
	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(subcommands.FlagsCommand(), "")
	subcommands.Register(subcommands.CommandsCommand(), "")
	subcommands.Register(&lastRunCmd{}, "")
	subcommands.Register(&preRunCmd{}, "")
	subcommands.Register(&postRunCmd{}, "")
	flag.Parse()
	ctx := context.Background()
	return subcommands.Execute(ctx)
}
