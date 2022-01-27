package pkg

import (
	"errors"

	"github.com/bitfield/script"
)

func CheckDir() (bool, error) {
	command := "git rev-parse --is-inside-work-tree"
	for _, c := range []string{command} {
		p := script.Exec(c)
		if err := p.Error(); err != nil {
			p.SetError(nil)
			output, _ := p.String()
			return false, errors.New(output)
		} else {
			output, _ := p.String()
			if output == "true" {
				return true, nil
			}
		}
	}
	return true, nil
}
