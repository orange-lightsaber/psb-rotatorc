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

type postRunCmd struct {
	Run
}

func (*postRunCmd) Name() string     { return "postrun" }
func (*postRunCmd) Synopsis() string { return "Start post-run." }
func (*postRunCmd) Usage() string {
	return `postrun [-name] <string> [-compkey] <string> [-freq] <int> [-delay] <int> [-year] <int> [-month] <int> [-day] <int> [-initial] <int>:
  Starts post-run operation.
`
}

func (cmd *postRunCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	if ok := cmd.CheckFlags(); !ok {
		fmt.Print("invalid arguments")
		os.Exit(2)
	}
	req := sockets.Request{
		sockets.PostRun_Req,
		rotator.RunConfigData{
			CompatibilityKey: cmd.compkey,
			Name:             cmd.name,
			Frequency:        cmd.freq,
			RotationDelay:    cmd.delay,
			Year:             rotator.Year{Duration: cmd.year},
			Month:            rotator.Month{Duration: cmd.month},
			Day:              rotator.Day{Duration: cmd.day},
			Initial:          rotator.Initial{Duration: cmd.initial},
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
	// fmt.Print(res.Response)
	return subcommands.ExitSuccess
}
