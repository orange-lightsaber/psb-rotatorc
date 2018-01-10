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

type rotateCmd struct {
	name string
}

func (cmd *rotateCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&cmd.name, "name", "", "")
}

func (cmd *rotateCmd) CheckFlags() (r bool) {
	switch {
	case cmd.name == "":
	default:
		r = true
	}
	return
}

func (*rotateCmd) Name() string     { return "rotate" }
func (*rotateCmd) Synopsis() string { return "Start rotation." }
func (*rotateCmd) Usage() string {
	return `rotate [-name]:
  Starts rotation, prints UTC time of rotation.
`
}

func (cmd *rotateCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	if ok := cmd.CheckFlags(); !ok {
		fmt.Print("invalid arguments")
		os.Exit(2)
	}
	req := sockets.Request{
		sockets.Rotate_Req,
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
