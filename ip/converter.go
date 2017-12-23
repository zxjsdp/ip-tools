package ip

import (
	"regexp"
	"strings"

	"github.com/zxjsdp/gotool/gotool"
)

func PrepareInputString(source string) string {
	if len(source) == 0 {
		return ""
	}

	return gotool.ReplaceRunes(source, '\n', ' ')
}

func GetRange(ips []string) []string {
	if len(ips) == 0 {
		return ips
	}

	result := make([]string, 0)
	for _, ip := range ips {
		elements := strings.Split(ip, ".")
		firstPart := strings.Join(elements[:3], ".")
		result = append(result, firstPart+".0-255")
	}

	return gotool.RemoveDuplicates(result)
}

func GetSingleIPsByRegexp(source string) []string {
	ipRegexp, _ := regexp.Compile("\\b\\d+\\.\\d+\\.\\d+\\.\\d+\\b")
	ips := ipRegexp.FindAllString(source, -1)

	return gotool.RemoveDuplicates(ips)
}
