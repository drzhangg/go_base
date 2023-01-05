package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"k8s.io/kubectl/pkg/util/templates"
	"os"
)

func main() {

	var flag bool


	//cmd.NewCNSPCommand("dspctl", os.Stdin, os.Stdout, os.Stderr)

	cmds := &cobra.Command{
		Use: "jerry",
		Long: templates.LongDesc(`
	Cloud Native Service Platform Client

	This client helps you deploy and run your applications on the Cloud Native Service Platform`),
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	cmds.Flags().BoolVarP(&flag,"flag","f",flag,"test flag value")

	fmt.Println("falg:::",flag)

	cmds.AddCommand()


	if err := cmds.Execute(); err != nil {
		fmt.Printf("Error: %+v", err)
		os.Exit(1)
	}
	os.Exit(0)


}
