package cmd

import (
	"github.com/spf13/cobra"
	"github.com/pandazhuzi/buns/application"
)

func CmdPackage(a app.IApplication, ctx *Context) *cobra.Command {

	type options struct {
		out string
		project string
	}

	opts := new(options)

	cmd := &cobra.Command{
		Use:"package",
		Short:"package command",
		Run:func(cmd *cobra.Command, args []string) {
			println("package, args :", args)
		}}

	//TODO: add you flags here
	flags := cmd.Flags()
	flags.StringVarP(&opts.out,"out","o",a.GetCwd(),)


	return cmd
}