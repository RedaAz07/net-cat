package validators

import (
	"chat-app/utils"
)
// validators for usernames
func ValidateLength(username string) bool {
	if len(username) < 3 || len(username) > 15 {
		return false
	}
	return true
}

func ValidName(username string) bool {
	for _, i := range username {
		if !(i >= 'A' && i <= 'Z') && !(i >= 'a' && i <= 'z') {
			return false
		}
	}
	return true
}

func SameName(username string) bool {
	for _, client := range utils.Clients {
		if username == client {
			return false
		}
	}
	return true
}
// validators for meessages 
func ValidateLengthMessage(message string) bool {
	if len(message) < 1 || len(message) > 50 {
		return false
	}
	return true
}

func ValidMessage(message string) bool {
	for _, i := range message {
		if i < 32 || i > 126 {
			return false
		}
	}
	return true
}
