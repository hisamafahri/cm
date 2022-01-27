package data

type commitStruct struct {
	Name        string
	Description string
}

var CommitTypes = []commitStruct{
	{Name: "feat", Description: "A new feature"},
	{Name: "fix", Description: "A bug fix"},
	{Name: "docs", Description: "Documentation only changes"},
	{Name: "style", Description: "Changes that do not affect the meaning of the code (white-space, formatting, missing semi-colons, etc)"},
	{Name: "refactor", Description: "A code change that neither fixes a bug nor adds a feature"},
	{Name: "perf", Description: "A code change that improves performance"},
	{Name: "test", Description: "Adding missing or correcting existing tests"},
	{Name: "chore", Description: "Changes to the build process or auxiliary tools and libraries such as documentation generation"},
}
