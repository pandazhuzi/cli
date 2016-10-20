package cmd

import (
	"github.com/spf13/cobra"
	"github.com/pandazhuzi/core/project"
	"github.com/pandazhuzi/buns/application"
)

func CmdAdd(a app.IApplication, ctx *Context) *cobra.Command {

	type options struct {
		name string
		dir string
	}

	opts := new(options)

	cmd := &cobra.Command{
		Use:"add",
		Short:"add cli command",
		Run:func(cmd *cobra.Command, args []string) {

			if(len(opts.name) == 0){
				println("--name required")
				return
			}

			cli, err := core.OpenCliProject(opts.dir)

			if(err != nil){
				println(err.Error())
				return
			}

			err = cli.Add(opts.name,a.GetResource(),opts.dir)

			if(err != nil){
				println(err.Error())
				return
			}

		}}

	flags := cmd.Flags()
	flags.StringVarP(&opts.dir,"dir","d",a.GetCwd(),"cli project root")
	flags.StringVarP(&opts.name,"name","n","","command name")



	return cmd
}