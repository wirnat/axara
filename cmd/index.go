/*
	Copyright © aksara-tech 2022
*/
package cmd

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/wirnat/axara/cmd/v1/files"
	"github.com/wirnat/axara/cmd/v1/global"
	v1 "github.com/wirnat/axara/cmd/v1/runner"
	"github.com/wirnat/axara/cmd/v1/service/getter/go_git"
	"os"
	"text/template"
)

var generatorCmd = &cobra.Command{
	Use:     "generate",
	Short:   "Auto generate Design Pattern",
	Example: "axara generate conf.yaml --models User",
	Long:    `Some folks say that Design Patterns are dead. How foolish. The Design Patterns book is one of the most important books published in our industry.  The concepts within should be common rudimentary knowledge for all professional programmers.`,
	Run: func(cmd *cobra.Command, args []string) {
		app := v1.NewApp(args[0])
		err := app.Generate()
		if err != nil {
			logrus.Fatal(err)
		}
	},
}

var checkVersion = &cobra.Command{
	Use:   "version",
	Short: "Check Version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Axara version " + VERSION)
	},
}

var newConfig = &cobra.Command{
	Use:   "new",
	Short: "New Axara Config file",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("❌ require argument filename, example: axara new wirnat_arch")
			return
		}
		file, err := os.Create(args[0] + ".yaml")
		if err != nil {
			fmt.Println("❌ " + err.Error())
			return
		}
		tmt, err := template.New("new config").Parse(files.New)
		if err != nil {
			fmt.Println("❌ " + err.Error())
			return
		}

		err = tmt.Execute(file, nil)
		if err != nil {
			fmt.Println("❌ " + err.Error())
			return
		}
	},
}

var getter = &cobra.Command{
	Use:        "get",
	Aliases:    nil,
	SuggestFor: nil,
	Short:      "Get CLI Item from github",
	Example:    "axara get https://github.com/wirnat/axara-template-go-clean-architecture templates",
	Run: func(cmd *cobra.Command, args []string) {
		gt := go_git.NewGetter()
		if len(args) != 2 {
			logrus.Fatal("directory required")
		}
		err := gt.Get(args[0], args[1])
		if err != nil {
			logrus.Fatal(err)
		}
	},
	RunE:                       nil,
	PostRun:                    nil,
	PostRunE:                   nil,
	PersistentPostRun:          nil,
	PersistentPostRunE:         nil,
	FParseErrWhitelist:         cobra.FParseErrWhitelist{},
	CompletionOptions:          cobra.CompletionOptions{},
	TraverseChildren:           false,
	Hidden:                     false,
	SilenceErrors:              false,
	SilenceUsage:               false,
	DisableFlagParsing:         false,
	DisableAutoGenTag:          false,
	DisableFlagsInUseLine:      false,
	DisableSuggestions:         false,
	SuggestionsMinimumDistance: 0,
}

func init() {
	rootCmd.AddCommand(generatorCmd)
	rootCmd.AddCommand(checkVersion)
	rootCmd.AddCommand(getter)
	rootCmd.AddCommand(newConfig)
	rootCmd.PersistentFlags().StringSliceVarP(&global.Tags, "tags", "g", []string{}, "List of execute traits/jobs")
	rootCmd.PersistentFlags().StringSliceVarP(&global.ExecuteModels, "models", "m", []string{}, "list of execute models")
}
