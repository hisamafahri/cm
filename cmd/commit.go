package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/bitfield/script"
	"github.com/logrusorgru/aurora"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var All bool

type commitStruct struct {
	Name        string
	Description string
}

var commitTypes = []commitStruct{
	{Name: "feat", Description: "A new feature"},
	{Name: "fix", Description: "A bug fix"},
	{Name: "docs", Description: "Documentation only changes"},
	{Name: "style", Description: "Changes that do not affect the meaning of the code (white-space, formatting, missing semi-colons, etc)"},
	{Name: "refactor", Description: "A code change that neither fixes a bug nor adds a feature"},
	{Name: "perf", Description: "A code change that improves performance"},
	{Name: "test", Description: "Adding missing or correcting existing tests"},
	{Name: "chore", Description: "Changes to the build process or auxiliary tools and libraries such as documentation generation"},
}

var templates = &promptui.SelectTemplates{
	Label:    "{{ . }}?",
	Active:   "> {{ .Name | cyan }}: {{ .Description | red }}",
	Inactive: "  {{ .Name | cyan }}: {{ .Description | red }}",
	Selected: "> {{ .Name | cyan }}: {{ .Description | red }}",
	// 		Details: `
	// --------- Change Type ----------
	// {{ "Name:" | faint }}	{{ .Name }}
	// {{ "Desc:" | faint }}	{{ .Description }}`,
}

var commitCommand = &cobra.Command{
	Use:   "commit",
	Short: "Commit is CLI to help you structurized your commit message",
	Long: `Simple but powerful CLI to help your commit message to follow
		conventional commit message`,
	Run: func(cmd *cobra.Command, args []string) {

		allFlag, _ := cmd.Flags().GetBool("all")

		changeTypePrompt := promptui.Select{
			Label:     "Type of Change",
			Items:     commitTypes,
			Templates: templates,
			Size:      8,
		}

		// Scope

		validateScope := func(input string) error {
			if len(input) > 25 {
				return errors.New("commit scope must have less than 25 characters")
			} else if len(input) < 2 {
				return errors.New("commit scope must have more than 2 characters")
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
			} else if len(input) < 5 {
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
			fmt.Printf("Commit type is required: %v\n", errType)
			return
		}

		commitScope, errScope := scopePrompt.Run()

		if errScope != nil {
			fmt.Printf("Commit failed %v\n", errScope)
			return
		}

		commitMessage, errMessage := commitMessagePrompt.Run()

		if errMessage != nil {
			fmt.Printf("Commit failed %v\n", errMessage)
			return
		}

		fullCommit := "git commit -m \"" + commitTypes[i].Name + "(" + commitScope + "): " + commitMessage + "\""

		if allFlag {
			for _, c := range []string{"git add ."} {
				p := script.Exec(c)
				// fmt.Println("Exit Status:", p.ExitStatus())
				if err := p.Error(); err != nil {
					p.SetError(nil)
					out, _ := p.Stdout()
					fmt.Println(out)
				} else {
					fmt.Println(aurora.Black(" Status OK ").BgGreen().Bold(), "Successfully staged changes")
				}
			}
		}

		for _, c := range []string{fullCommit} {
			p := script.Exec(c)
			// fmt.Println("Exit Status:", p.ExitStatus())
			if err := p.Error(); err != nil {
				p.SetError(nil)
				out, _ := p.Stdout()
				fmt.Println(out)
			} else {
				fmt.Println(aurora.Black(" Processing ").BgBrightWhite().Bold())
				p.Stdout()
				fmt.Println(aurora.Black(" Status OK ").BgGreen().Bold(), "Successfully commit changes")
			}
		}
	},
}

func Execute() {
	if err := commitCommand.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	commitCommand.PersistentFlags().BoolVarP(&All, "all", "a", false, "Commit all changes in current directory")
}
