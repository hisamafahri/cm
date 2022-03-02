package cmd

import (
	"fmt"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/hisamafahri/cm/helper"
	"github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
)

var pushCommand = &cobra.Command{
	Use:     "push",
	Aliases: []string{"p"},
	Short:   "Push to remote repo",
	Long:    `Push to remote repo, by default on the current branch and your chosen remote repo`,
	Run: func(cmd *cobra.Command, args []string) {
		// Check if current directory is a git repository
		// if it is not, return an error
		_, err := helper.CheckDir()

		if err != nil {
			fmt.Println(aurora.White(" ERROR ").BgRed().Bold(), err.Error())
			return
		}

		// Check the current branch name
		// Return error if there is an error
		branch, err := helper.GetBranch()

		if err != nil {
			fmt.Println(aurora.White(" ERROR ").BgRed().Bold(), err.Error())
			return
		}

		// Get all of the remote repository and aliases
		// Return error if there is an error
		remotes, err := helper.GetRemote()

		if err != nil {
			fmt.Println(aurora.White(" ERROR ").BgRed().Bold(), err.Error())
			return
		}

		// If there is only ONE remote repo, just push the current branch right away
		if len(remotes) == 1 {
			remoteAlias := strings.Split(remotes[0], ": ")
			pushCommand := "git push " + strings.TrimSuffix(remoteAlias[0], "\n") + " " + branch
			helper.Push(pushCommand)
			return
		}

		// If there is multiple remote repo, ask the user to choose
		// Struct for chosen remote alias
		chosen := struct {
			Alias string `survey:"alias"`
		}{}

		// construct the prompt
		var chooseAliases = []*survey.Question{
			{
				Name: "alias",
				Prompt: &survey.Select{
					Message: "Which repository you want to push?:",
					Options: remotes,
					Default: remotes[0],
				},
			},
		}

		// perform the questions
		// and return error if there is a problem
		errPrompt := survey.Ask(chooseAliases, &chosen)

		if errPrompt != nil {
			fmt.Println(aurora.White(" ERROR ").BgRed().Bold(), "Push failed!")
			return
		}

		// get the alias name only
		remoteAlias := strings.Split(chosen.Alias, ": ")

		// formulate the push command
		pushCommand := "git push " + strings.TrimSuffix(remoteAlias[0], "\n") + " " + branch

		// push it
		helper.Push(pushCommand)
	},
}
