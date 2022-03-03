package helper

import (
	"errors"

	"github.com/bitfield/script"
)

func GetLog() (string, error) {
	var output string

	command := "git log --pretty=\"format:\x1B[38;2;249;38;114m%h\x1B[0m \u001b[33m%d\u001b[0m %s \u001b[38;5;242m %cn - %cr\u001b[0m  \""
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
