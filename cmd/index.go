/*
Copyright © 2022
*/
package cmd

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var generatorCmd = &cobra.Command{
	Use:   "generate",
	Short: "Auto generate Design Pattern",
	Long:  `Some folks say that Design Patterns are dead. How foolish. The Design Patterns book is one of the most important books published in our industry.  The concepts within should be common rudimentary knowledge for all professional programmers.`,
	Run: func(cmd *cobra.Command, args []string) {
		conf, err := OpenConfig(args[0])
		if err != nil {
			logrus.Fatalln(err.Error())
		}

		if conf.Key == "ᬅᬓ᭄ᬱᬭ" {
			reader := NewModelScanner(conf)
			blueprint, err := reader.Scan()
			if err != nil {
				logrus.Fatalln(err.Error())
			}

			decoder := NewDecoder(blueprint)
			blueprint, err = decoder.Decode()
			if err != nil {
				logrus.Fatalln(err.Error())
			}

			err = blueprint.Generate()
			if err != nil {
				logrus.Fatalln(err.Error())
			}
		} else {
			logrus.Fatalln("invalid key in json file")
		}
	},
}

const VERSION = "v0.0.2"

var checkVersion = &cobra.Command{
	Use:   "version",
	Short: "Check CLI Version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Aksara CLI version " + VERSION)
	},
}

func init() {
	rootCmd.AddCommand(generatorCmd)
	rootCmd.AddCommand(checkVersion)
}
