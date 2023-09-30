package utils

import (
	"regexp"
	"strings"
)

type Home struct {
	Username string
	System string
	Greetings string
	Length string
	Filename string
	Errors map[string]string
}

func (msg *Home) Validate() bool {
	msg.Errors = make(map[string]string)

	if strings.TrimSpace(msg.Username) == "" {
		msg.Errors["Username"] = "Please enter a username"
	}

	if strings.TrimSpace(msg.System) == "" {
		msg.Errors["System"] = "Please enter a system"
	}

	if !isNumeric(msg.Length) {
		msg.Errors["Length"] = "Please enter a number"
	}

	if strings.TrimSpace(msg.Greetings) == "" {
		msg.Errors["Greetings"] = "Please enter a greetings"
	}

	if strings.TrimSpace(msg.Filename) == "" {
		msg.Errors["Filename"] = "Please enter a filename"
	}

	return len(msg.Errors) == 0
}

func isNumeric(s string) bool {
  re := regexp.MustCompile(`^\d+$`)
  return re.MatchString(s)
}