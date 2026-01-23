package flagparser

import "errors"

func ParseArgs(args []string) error {
	if len(args) < 2 {
		return errors.New("error")
	}
	if len(args) > 5 {
		return errors.New("error")
	}
	return nil
}
