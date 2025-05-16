package fnCase

import "strings"

// CamelCase
// 영문문자열을 카멜케이스로 바꿔주는 함수
func CamelCase(v string) string {
	return camelCase(v, false)
}

// PascalCase
// 영문문자열을 파스칼케이스로 바꿔주는 함수
func PascalCase(v string) string {
	return camelCase(v, true)
}

func camelCase(s string, startUpperCase bool) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return s
	}

	n := strings.Builder{}
	n.Grow(len(s))
	capNext := startUpperCase
	for i, v := range []byte(s) {
		vIsCap := v >= 'A' && v <= 'Z'
		vIsLow := v >= 'a' && v <= 'z'
		if capNext {
			if vIsLow {
				v += 'A'
				v -= 'a'
			}
		} else if i == 0 {
			if vIsCap {
				v += 'a'
				v -= 'A'
			}
		}
		if vIsCap || vIsLow {
			n.WriteByte(v)
			capNext = false
		} else if vIsNum := v >= '0' && v <= '9'; vIsNum {
			n.WriteByte(v)
			capNext = true
		} else {
			capNext = v == '_' || v == ' ' || v == '-' || v == '.'
		}
	}
	return n.String()
}
