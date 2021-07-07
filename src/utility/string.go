package utility

import (
	"regexp"
)

// RemoveMultiSpaces : 스페이스 삭제
func RemoveMultiSpaces(str string) string {
	reLeadcloseWhtsp := regexp.MustCompile(`^[\s\p{Zs}]+|[\s\p{Zs}]+$`)
	reInsideWhtsp := regexp.MustCompile(`[\s\p{Zs}]{2,}`)
	result := reLeadcloseWhtsp.ReplaceAllString(str, "")
	result = reInsideWhtsp.ReplaceAllString(result, " ")
	return result
}
