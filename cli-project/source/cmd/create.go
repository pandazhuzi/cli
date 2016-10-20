package cmd

import (
	"github.com/spf13/cobra"
	"github.com/pandazhuzi/core/project"
	"github.com/pandazhuzi/buns/application"
)

func CmdCreate(a app.IApplication, ctx *Context) *cobra.Command {

	type options struct {
		name string
		out string
	}

	opts := new(options)

	cmd := &cobra.Command{
		Use:"create",
		Short:"create cli project",
		Run:func(cmd *cobra.Command, args []string) {

			if(len(opts.name) == 0){
				println("--name required")
				return
			}

			_, err := core.CreateCliProject(opts.name,a.GetResource(),opts.out)

			if(err != nil){
				println(err.Error())
				return
			}

		}}

	flags := cmd.Flags()
	flags.StringVar(&opts.name,"name","","set project name")
	flags.StringVar(&opts.out, "out",a.GetCwd(),"generate output dir")

	return cmd
}
