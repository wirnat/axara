/*
	Copyright © aksara-tech 2022
*/
package cmd

import (
	"fmt"
	"github.com/dgraph-io/badger"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	v1 "github.com/wirnat/axara/cmd/v1"
	"github.com/wirnat/axara/cmd/v1/key"
	"log"
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

const VERSION = "v1.0.7"

var checkVersion = &cobra.Command{
	Use:   "version",
	Short: "Check Version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Axara version " + VERSION)
	},
}

var setter = &cobra.Command{
	Use:   "set",
	Short: "Set Configuration",
	Long:  "--git-key ",
	Run: func(cmd *cobra.Command, args []string) {
		gitkey, err := cmd.Flags().GetString(key.GitKey)
		if err != nil {
			logrus.Fatal(err)
		} else {
			opts := badger.DefaultOptions(key.Storage())
			opts.Logger = nil

			db, err := badger.Open(opts)
			if err != nil {
				log.Fatal(err)
			}
			defer db.Close()

			db.Update(func(txn *badger.Txn) error {
				return txn.Set([]byte(key.GitKey), []byte(gitkey))
			})
		}
	},
}

var getter = &cobra.Command{
	Use:   "get",
	Short: "Get CLI Item from github",
	Run: func(cmd *cobra.Command, args []string) {
		gp := v1.NewGitPuller()
		if len(args) != 2 {
			logrus.Fatalf("invalid get argument, ex: axara get github.com/wirnat/basic-template template")
		}
		err := gp.Pull(args[0], args[1])
		if err != nil {
			logrus.Fatal(err)
		}
	},
}

func init() {
	rootCmd.PersistentFlags().String("git-key", "", "set git access token, ex: set --git-key your-token")
	rootCmd.AddCommand(generatorCmd)
	rootCmd.AddCommand(checkVersion)
	rootCmd.AddCommand(getter)
	rootCmd.AddCommand(setter)
}
