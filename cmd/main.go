package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/hisamafahri/commit/helper"
	"github.com/hisamafahri/commit/src/base"
	"github.com/hisamafahri/commit/src/functions"
	"github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
)

var All bool

var commitCommand = &cobra.Command{
	Use:   "commit",
	Short: "Commit is CLI to help you structurized your commit message",
	Long: `Simple but powerful CLI to help your commit message to follow
		conventional commit message`,
	Run: func(cmd *cobra.Command, args []string) {
		/*
			Check if current directory is a git repository
			if it is not, return an error
		*/
		_, err := helper.CheckDir()

		if err != nil {
			fmt.Println(aurora.White(" ERROR ").BgRed().Bold(), err.Error())
			return
		}

		/*
			Run the prompt or return an error
		*/

		result, err := functions.RunPrompt()

		if err != nil {
			fmt.Println(aurora.White(" ERROR ").BgRed().Bold(), "Commit failed!")
			return
		}

		/*
			Formulating the commit message
		*/

		index, _ := strconv.Atoi(result[0])

		commitCommand := "git commit -m \"" + base.CommitTypes[index].Name + "(" + result[1] + "): " + result[2] + "\""

		/*
			Get the -a or --all flag
		*/
		allFlag, _ := cmd.Flags().GetBool("all")

		/*
			if flag -a or --all exist, run `git add .` command
		*/
		if allFlag {
			addAllChanges()
		}

		/*
			Commit the changes
		*/
		commit(commitCommand)
	},
}

func Execute() {
	if err := commitCommand.Execute(); err != nil {
		fmt.Println(aurora.White(" ERROR ").BgRed().Bold(), err)
		os.Exit(1)
	}
}

func init() {
	commitCommand.PersistentFlags().BoolVarP(&All, "all", "a", false, "Commit all changes in current directory")
}
