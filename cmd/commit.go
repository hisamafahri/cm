package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/bitfield/script"
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

		validateScope := func(input string) error {
			if len(input) > 25 {
				return errors.New("commit scope must have less than 25 characters")
			} else if len(input) < 2 {
				return errors.New("commit scope must have more than 1 characters")
			}
			return nil
		}

		scopePrompt := promptui.Prompt{
			Label:    "Scope of changes (eg. file, function, etc)",
			Validate: validateScope,
		}

		// message

		validateCommitMessage := func(input string) error {
			if len(input) > 100 {
				return errors.New("commit message must have less than 100 characters")
			} else if len(input) < 6 {
				return errors.New("commit message must have more than 5 characters")
			}
			return nil
		}

		commitMessagePrompt := promptui.Prompt{
			Label:    "Commit message title (min 5 & max 100)",
			Validate: validateCommitMessage,
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

		fullCommit := "git commit -m \"" + data.CommitTypes[i].Name + "(" + commitScope + "): " + commitMessage + "\""

		if allFlag {
			for _, c := range []string{"git add ."} {
				fmt.Println(aurora.Black(" INFO ").BgBrightWhite().Bold(), "Staging changes...")
				p := script.Exec(c)
				if err := p.Error(); err != nil {
					p.SetError(nil)
					output, _ := p.String()
					fmt.Println(aurora.White(" ERROR ").BgRed().Bold(), output)
				} else {
					fmt.Println(aurora.Black(" SUCCESS ").BgGreen().Bold(), "Successfully staged changes")
				}
			}
		}

		for _, c := range []string{fullCommit} {
			fmt.Println() // add break line
			fmt.Println(aurora.Black(" INFO ").BgBrightWhite().Bold(), "Committing changes...")
			p := script.Exec(c)
			if err := p.Error(); err != nil {
				p.SetError(nil)
				output, _ := p.String()
				fmt.Println(aurora.White(" ERROR ").BgRed().Bold(), output)
			} else {
				output, _ := p.String()
				fmt.Println(aurora.Black(" INFO ").BgBrightWhite().Bold(), output)
				fmt.Println(aurora.Black(" SUCCESS ").BgGreen().Bold(), "Successfully commit changes")
			}
		}
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
