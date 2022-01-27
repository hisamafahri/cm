package functions

import "github.com/manifoldco/promptui"

var Templates = &promptui.SelectTemplates{
	Label:    "{{ . }}?",
	Active:   "> {{ .Name | cyan }}: {{ .Description | red }}",
	Inactive: "  {{ .Name | cyan }}: {{ .Description | red }}",
	Selected: "> {{ .Name | cyan }}: {{ .Description | red }}",
}
