package common

import (
	"strings"
)

func TransformError(errString string) []string {
	errors := []string{}
	splitted := strings.Split(errString, "\n")
	for _, v := range splitted {
		split := strings.Split(v, "' Error:")
		err := split[0]
		if len(split) > 1 {
			err = split[1]
		}
		errors = append(errors, err)
	}
	return errors
}
