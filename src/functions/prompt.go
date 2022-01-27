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

	i, _, errT := changeTypePrompt.Run()

	commitScope, errS := scopePrompt.Run()

	commitMessage, errM := commitMessagePrompt.Run()

	/*
		Need better error handling approach
	*/

	if errT != nil {
		return [3]string{}, errT
	} else if errS != nil {
		return [3]string{}, errS
	} else if errM != nil {
		return [3]string{}, errM
	}

	return [3]string{strconv.Itoa(i), commitScope, commitMessage}, nil

}
