package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)

	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[=\-\*~]*>`)

	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`(?i)".*password.*"`)

	var cnt int = 0
	for _, l := range lines {
		if re.MatchString(l) {
			cnt++
		}
	}
	return cnt
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d+`)

	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User\s+([a-zA-Z\d]+)`)
	for i, l := range lines {
		matches := re.FindStringSubmatch(l)
		if len(matches) > 1 {
			lines[i] = fmt.Sprintf("%s %s %s", "[USR]", matches[1], l)
		}
	}
	return lines
}
