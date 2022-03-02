package helper

import (
	"errors"
	"strings"

	"github.com/bitfield/script"
)

func GetRemote() ([]string, error) {
	remoteAlias, err := getRemoteAlias()
	if err != nil {
		return []string{}, err
	}

	remoteAliasList := strings.Fields(remoteAlias)

	var output []string
	for _, alias := range remoteAliasList {
		command := "git remote get-url --push " + alias
		for _, c := range []string{command} {
			p := script.Exec(c)
			if err := p.Error(); err != nil {
				p.SetError(nil)
				output, _ := p.String()
				return []string{}, errors.New(output)
			} else {
				result, _ := p.String()
				output = append(output, strings.TrimSuffix(alias, "\n")+": "+strings.TrimSuffix(result, "\n"))
			}
		}
	}
	return output, nil
}

func getRemoteAlias() (string, error) {
	command := "git remote"
	var output string
	for _, c := range []string{command} {
		p := script.Exec(c)
		if err := p.Error(); err != nil {
			p.SetError(nil)
			output, _ := p.String()
			return "", errors.New(output)
		} else {
			result, _ := p.String()
			output = result
		}
	}
	return output, nil
}
