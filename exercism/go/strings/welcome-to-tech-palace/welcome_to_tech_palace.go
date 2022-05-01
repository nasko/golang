package techpalace

import (
	"fmt"
	"strings"
)

func WelcomeMessage(customer string) string {
	return fmt.Sprintf("Welcome to the Tech Palace, %v", strings.ToUpper(customer))
}

func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	nl := string('\n')
	border := ""

	for i := 0; i < numStarsPerLine; i++ {
		border += string('*')
	}
	return border + nl + welcomeMsg + nl + border
}

func CleanupMessage(oldMsg string) string {
	sanitized := strings.ReplaceAll(oldMsg, "*", "")
	sanitized = strings.ReplaceAll(sanitized, "\r\n", "")
	sanitized = strings.ReplaceAll(sanitized, "\n", "")
	sanitized = strings.TrimSpace(sanitized)
	return sanitized
}
