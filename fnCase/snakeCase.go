package fnCase

import (
	"regexp"
	"strings"
)

// SnakeCase
// 영문문자열을 스네이크케이스로 바꿔주는 함수
func SnakeCase(s string) string {
	matchFirstCap := regexp.MustCompile("(.)([A-Z][a-z]+)")
	matchAllCap := regexp.MustCompile("([a-z0-9])([A-Z])")
	snake := matchFirstCap.ReplaceAllString(s, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}
