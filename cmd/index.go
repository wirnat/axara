/*
Copyright © 2022
*/
package cmd

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	v1 "github.com/wirnat/axara/cmd/v1"
)

var generatorCmd = &cobra.Command{
	Use:   "generate",
	Short: "Auto generate Design Pattern",
	Long:  `Some folks say that Design Patterns are dead. How foolish. The Design Patterns book is one of the most important books published in our industry.  The concepts within should be common rudimentary knowledge for all professional programmers.`,
	Run: func(cmd *cobra.Command, args []string) {
		app := v1.NewApp(args[0])
		err := app.Generate()
		if err != nil {
			logrus.Fatal(err)
		}
	},
}

const VERSION = "v2.0.0"

var checkVersion = &cobra.Command{
	Use:   "version",
	Short: "Check CLI Version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Aksara CLI version " + VERSION)
	},
}
var getter = &cobra.Command{
	Use:   "get",
	Short: "Get CLI Item from github",
	Run: func(cmd *cobra.Command, args []string) {
		gp := v1.NewGitPuller()
		if len(args) != 2 {
			logrus.Fatalf("invalid get argument, ex: aksara-cli get github.com/wirnat/basic-template template")
		}
		err := gp.Pull(args[0], args[1])
		if err != nil {
			logrus.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(generatorCmd)
	rootCmd.AddCommand(checkVersion)
	rootCmd.AddCommand(getter)
}
