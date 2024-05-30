package parsinglogfiles

import (
	"fmt"
	"regexp"
	"strings"
)

func IsValidLine(text string) bool {
	validPrefixes := []string{"[TRC]", "[DBG]", "[INF]", "[WRN]", "[ERR]", "[FTL]"}
	for _, prefix := range validPrefixes {
		if strings.HasPrefix(text, prefix) {
			return true
		}
	}
	return false
}

func SplitLogLine(text string) []string {
	separatorRegex := `<[~*=\-]*>`
	re := regexp.MustCompile(separatorRegex)
	res := re.Split(text, -1)
	return res
}

func CountQuotedPasswords(lines []string) int {
	count := 0
	regex := `(?i)"[^"]*password[^"]*"`
	re := regexp.MustCompile(regex)
	for _, line := range lines {
		match := re.MatchString(line)
		if match {
			count += 1
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	regex := `end-of-line\d+`
	re := regexp.MustCompile(regex)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	pattern := `User\s+(\S+)`
	re := regexp.MustCompile(pattern)

	var result []string
	for _, log := range lines {
		// Check if the log line contains "User " followed by a user name
		if re.MatchString(log) {
			// Extract the user name
			userName := re.FindStringSubmatch(log)[1]
			// Prefix the log line with "[USR] userName "
			taggedLog := fmt.Sprintf("[USR] %s %s", userName, log)
			result = append(result, taggedLog)
		} else {
			// If no match, add the log line unchanged
			result = append(result, log)
		}
	}

	return result
}
