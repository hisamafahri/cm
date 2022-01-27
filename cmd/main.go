package cmd

import (
	"fmt"
	"os"

	"github.com/hisamafahri/commit/data"
	"github.com/hisamafahri/commit/helper"
	"github.com/logrusorgru/aurora"
	"github.com/manifoldco/promptui"
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

		allFlag, _ := cmd.Flags().GetBool("all")

		changeTypePrompt := promptui.Select{
			Label:     "Type of Change",
			Items:     data.CommitTypes,
			Templates: data.Templates,
			Size:      8,
		}

		// Scope

		scopePrompt := promptui.Prompt{
			Label:    "Scope of changes (eg. file, function, etc)",
			Validate: data.ValidateScope,
		}

		// message

		commitMessagePrompt := promptui.Prompt{
			Label:    "Commit message title (min 5 & max 100)",
			Validate: data.ValidateCommitMessage,
		}

		i, _, errType := changeTypePrompt.Run()

		if errType != nil {
			fmt.Println(aurora.White(" ERROR ").BgRed().Bold(), "Commit type is required!")
			return
		}

		commitScope, errScope := scopePrompt.Run()

		if errScope != nil {
			fmt.Println(aurora.White(" ERROR ").BgRed().Bold(), "Commit failed!")
			return
		}

		commitMessage, errMessage := commitMessagePrompt.Run()

		if errMessage != nil {
			fmt.Println(aurora.White(" ERROR ").BgRed().Bold(), "Commit failed!")
			return
		}

		commitCommand := "git commit -m \"" + data.CommitTypes[i].Name + "(" + commitScope + "): " + commitMessage + "\""

		if allFlag {
			addAllChanges()
		}

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
