package functions

import "errors"

var ValidateScope = func(input string) error {
	if len(input) > 25 {
		return errors.New("commit scope must have less than 25 characters")
	} else if len(input) < 2 {
		return errors.New("commit scope must have more than 1 characters")
	}
	return nil
}

var ValidateCommitMessage = func(input string) error {
	if len(input) > 100 {
		return errors.New("commit message must have less than 100 characters")
	} else if len(input) < 6 {
		return errors.New("commit message must have more than 5 characters")
	}
	return nil
}
