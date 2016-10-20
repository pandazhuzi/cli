package main

import (
	"./cmd"
	"fmt"
	"github.com/pandazhuzi/buns/application"
	"github.com/spf13/cobra"
	"github.com/pandazhuzi/buns/errors"
)

func main() {

	cli := &app.CommandLine{
		Root: &cobra.Command{
			Use: "{{.Name}}",
			Long: fmt.Sprintf("{{.Name}} is a \n"), //TODO:Add description here
			Run: func(cmd *cobra.Command, args []string) {
				cmd.Help()
			}},
		Main: func(c *app.CommandLine) error{

			ctx, err := cmd.NewContext()

			if err != nil {
				return errors.MakeError(err)
			}

			//TODO: Add you commands here
			c.AddCommand(
				cmd.CmdVersion(c, ctx))

			return nil
		}}

	cli.Run()

}
