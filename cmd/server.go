/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/serverless-cloud-robot/robo9-cli/internal/controller"
	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Starts up web service for managing infra jobs",
	Long: `When deploying this in a container you can expect this service to run 
	on port 18080. You can override this port by specifying a -port command
	

You can see more details on how to use the cli by following our docs `,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("starting api web service. v0.0.1 ...")
		controller.StartServer()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
