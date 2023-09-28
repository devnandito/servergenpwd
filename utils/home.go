package utils

import "strings"

type ValidateHome struct {
	Username string
	System string
	Errors map[string]string
}

func (msg *ValidateHome) Validate() bool {
	msg.Errors = make(map[string]string)

	if strings.TrimSpace(msg.Username) == "" {
		msg.Errors["Username"] = "Please enter a username"
	}

	return len(msg.Errors) == 0
}