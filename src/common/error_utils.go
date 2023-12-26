package common

import (
	"strings"
)

func TransformError(errString string) []string {
	errors := []string{}
	splitted := strings.Split(errString, "\n")
	for _, v := range splitted {
		split := strings.Split(v, "' Error:")
		errors = append(errors, split[1])
	}
	return errors
}
