package cmd

import (
	"github.com/spf13/cobra"
	"github.com/pandazhuzi/buns/application"
)

//TODO: set you cli version here
const version = "0.0.1"

func CmdVersion(a app.IApplication, ctx *Context) *cobra.Command {

	cmd := &cobra.Command{
		Use:"version",
		Short:"show version",
		Run:func(cmd *cobra.Command, args []string) {
			println(version)
		}}

	return cmd
}
