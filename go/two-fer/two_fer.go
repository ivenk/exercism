package twofer

import (
	"strings"
	"fmt"
	)

const defaultValue = "you"

func ShareWith(name string) string {
	//check if the name is empty or just spaces
	if strings.TrimSpace(name) == "" {
		name = defaultValue
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
