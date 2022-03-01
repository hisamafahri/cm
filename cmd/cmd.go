package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/hisamafahri/cm/helper"
	"github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
)

var All bool

var commitCommand = &cobra.Command{
	Use:   "cm",
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

		// result, err := functions.RunPrompt()
		answers := struct {
			Type    string `survey:"types"`
			Scope   string
			Message string
		}{}

		// perform the questions
		errPrompt := survey.Ask(helper.Prompt, &answers)

		if errPrompt != nil {
			fmt.Println(aurora.White(" ERROR ").BgRed().Bold(), "Commit failed!")
			return
		}

		// Split the commit type string
		typeAbbrv := strings.Split(answers.Type, ":")

		/*
			Formulating the commit message
		*/

		commitCommand := "git commit -m \"" + typeAbbrv[0] + "(" + answers.Scope + "): " + answers.Message + "\""

		/*
			Get the -a or --all flag
		*/
		allFlag, _ := cmd.Flags().GetBool("all")

		/*
			if flag -a or --all exist, run `git add .` command
		*/
		if allFlag {
			helper.AddAllChanges()
		}

		/*
			Commit the changes
		*/
		helper.Commit(commitCommand)
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
