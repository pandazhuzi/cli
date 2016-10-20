package cmd

import (
	"github.com/spf13/cobra"
	"github.com/pandazhuzi/buns/application"
)

func Cmd{{.CamelName}}(a app.IApplication, ctx *Context) *cobra.Command {

	//TODO: initalize you options here
	//type options struct {
	//}

	//opts := new(options)

	cmd := &cobra.Command{
		Use:"{{.UnixName}}",
		Short:"{{.UnixName}} command",
		Run:func(cmd *cobra.Command, args []string) {
			println("{{.UnixName}}, args :", args)
		}}

	//TODO: add you flags here
	//flags := cmd.Flags()


	return cmd
}