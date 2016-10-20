package cmd

import (
	"github.com/spf13/cobra"
	"github.com/pandazhuzi/buns/application"
	"github.com/pandazhuzi/core/project"
)

func CmdBuild(a app.IApplication, ctx *Context) *cobra.Command {

	type options struct {
		dir string
	}

	opts := new(options)

	cmd := &cobra.Command{
		Use:"build",
		Short:"build go main file to build folder",
		Run:func(cmd *cobra.Command, args []string) {

			cli, err := core.OpenCliProject(opts.dir)

			if(err != nil){
				println(err.Error())
				return
			}

			err = cli.Build()

			if err != nil {
				println(err.Error())
				return
			}

		}}


	flags := cmd.Flags()
	flags.StringVarP(&opts.dir,"dir","d",a.GetCwd(),"set cli directory")


	return cmd
}