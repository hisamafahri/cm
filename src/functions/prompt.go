package functions

import (
	"strconv"

	"github.com/hisamafahri/commit/src/base"
	"github.com/manifoldco/promptui"
)

func RunPrompt() ([3]string, error) {
	changeTypePrompt := promptui.Select{
		Label:     "Type of Change",
		Items:     base.CommitTypes,
		Templates: Templates,
		Size:      8,
	}

	scopePrompt := promptui.Prompt{
		Label:    "Scope of changes (eg. file, function, etc)",
		Validate: ValidateScope,
	}

	commitMessagePrompt := promptui.Prompt{
		Label:    "Commit message title (min 5 & max 100)",
		Validate: ValidateCommitMessage,
	}

	/*
		Need better error handling approach
	*/
	i, _, err := changeTypePrompt.Run()
	if err != nil {
		return [3]string{}, err
	}

	commitScope, err := scopePrompt.Run()
	if err != nil {
		return [3]string{}, err
	}

	commitMessage, err := commitMessagePrompt.Run()
	if err != nil {
		return [3]string{}, err
	}

	return [3]string{strconv.Itoa(i), commitScope, commitMessage}, nil

}
