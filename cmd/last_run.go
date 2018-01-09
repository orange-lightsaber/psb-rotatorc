package cmd

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/google/subcommands"
	"github.com/orange-lightsaber/psb-rotatord/rotator"
	"github.com/orange-lightsaber/psb-rotatord/sockets"
)

type lastRunCmd struct {
	name string
}

func (cmd *lastRunCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&cmd.name, "name", "", "")
}

func (cmd *lastRunCmd) CheckFlags() (r bool) {
	switch {
	case cmd.name == "":
	default:
		r = true
	}
	return
}

func (*lastRunCmd) Name() string     { return "lastrun" }
func (*lastRunCmd) Synopsis() string { return "Print time since last run." }
func (*lastRunCmd) Usage() string {
	return `lastrun [-name] <string>:
Prints time since last run.
`
}

func (cmd *lastRunCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	if ok := cmd.CheckFlags(); !ok {
		fmt.Print("invalid arguments")
		os.Exit(2)
	}
	req := sockets.Request{
		sockets.LastRun_Req,
		rotator.RunConfigData{
			Name: cmd.name,
		}}
	res, err := req.NewRequest()
	if err != nil {
		fmt.Print(err)
		os.Exit(2)
	}
	if res.Error != "" {
		fmt.Print(res.Error)
		os.Exit(2)
	}
	fmt.Print(res.Response)
	return subcommands.ExitSuccess
}
