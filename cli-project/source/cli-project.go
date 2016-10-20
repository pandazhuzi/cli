package main

import (
	"./cmd"
	"fmt"
	"github.com/pandazhuzi/buns/application"
	"github.com/spf13/cobra"
	"github.com/pandazhuzi/buns/errors"
)

const version = "0.0.1"

func main() {

	cli := &app.CommandLine{
		Root: &cobra.Command{
			Use: "cli-project",
			Long: fmt.Sprintf("buns cli for the development of the cli folder\n"+
				"version %s", version),
			Run: func(cmd *cobra.Command, args []string) {
				cmd.Help()
			}},
		Main: func(c *app.CommandLine) error{

			ctx, err := cmd.NewContext()

			if err != nil {
				return errors.MakeError(err)
			}

			c.AddCommand(
				cmd.CmdCreate(c,ctx),
				cmd.CmdAdd(c,ctx),
				cmd.CmdBuild(c,ctx))

			return nil
		}}

	cli.Run()

}
