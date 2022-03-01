package helper

import (
	"errors"
	"path/filepath"

	"github.com/AlecAivazis/survey/v2"
)

var Prompt = []*survey.Question{
	{
		Name: "types",
		Prompt: &survey.Select{
			Message: "Type of Change?:",
			Options: []string{
				"feat: A new feature",
				"fix: A bug fix",
				"docs: Documentation only changes",
				"style: Changes that do not affect the meaning of the code (white-space, formatting, missing semi-colons, etc)",
				"refactor: A code change that neither fixes a bug nor adds a feature",
				"perf: A code change that improves performance",
				"test: Adding missing or correcting existing tests",
				"chore: Changes to the build process or auxiliary tools and libraries such as documentation generation",
			},
			Default: "feat: A new feature",
		},
	},
	{
		Name: "scope",
		Prompt: &survey.Input{
			Message: "Scope of changes (eg. file, function, etc):",
			Suggest: func(toComplete string) []string {
				files, _ := filepath.Glob(toComplete + "*")
				return files
			},
		},
		Validate: func(val interface{}) error {
			// since we are validating an Input, the assertion will always succeed
			if str, ok := val.(string); !ok || len(str) > 25 || len(str) < 2 {
				return errors.New("this response cannot be less than 2 characters and longer than 25 characters")
			}
			return nil
		},
	},
	{
		Name:   "message",
		Prompt: &survey.Input{Message: "Commit message:"},
		Validate: func(val interface{}) error {
			// since we are validating an Input, the assertion will always succeed
			if str, ok := val.(string); !ok || len(str) > 50 || len(str) < 2 {
				return errors.New("this response cannot be less than 2 characters and longer than 50 characters")
			}
			return nil
		},
	},
}
