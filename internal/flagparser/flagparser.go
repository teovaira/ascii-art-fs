package flagparser

import (
	"errors"
	"strings"
)

func ParseArgs(args []string) error {
	count := 0
	if len(args) < 2 {
		return errors.New("error")
	}
	if len(args) > 5 {
		return errors.New("error")
	}
	if err := validateColorFlag(args); err != nil {
		return err
	}
	for i := 0; i < len(args); i++ {
		if strings.HasPrefix(args[i], "--color=") {
			count++
			if count > 1 {
				return errors.New("error")
			}
		}
	}
	if strings.HasPrefix(args[1], "--color=") && len(args) < 3 {
		return errors.New("error")
	}
	if strings.HasPrefix(args[1], "--color=") {
		checkColorInTheFlag := strings.Split(args[1], "=")
		if len(checkColorInTheFlag) > 1 && checkColorInTheFlag[1] == "" {
			return errors.New("error")
		}
		color := checkColorInTheFlag[1]
		if color != "" {
			if err := validColors(color); err != nil {
				return err
			}
		}
	}
	return nil

}
func validateColorFlag(args []string) error {
	isItAFlag := strings.HasPrefix(args[1], "-")
	if isItAFlag == true {

		firstTwoLetters := strings.HasPrefix(args[1], "--")

		if firstTwoLetters == false {
			return errors.New("error")
		}

		equalExistance := strings.Contains(args[1], "=")
		if equalExistance == false {
			return errors.New("error")
		}

	}
	return nil
}
func validColors(color string) error {
	allowedColors := map[string]bool{
		"red":     true,
		"green":   true,
		"yellow":  true,
		"orange":  true,
		"blue":    true,
		"magenta": true,
	}
	_, exists := allowedColors[color]
	if !exists {
		return errors.New("error")
	}

	return nil
}
