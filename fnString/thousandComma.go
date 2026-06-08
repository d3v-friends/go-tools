package fnString

import (
	"bytes"
	"fmt"
	"strings"
)

// write by Gemini

// Number Go의 모든 정수 및 실수 타입을 아우르는 인터페이스 제약
type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

// ThousandComma 모든 숫자 타입을 받아 1000 단위 콤마를 찍어 반환합니다.
func ThousandComma[T Number](num T) string {
	// 1. float64로 정밀도를 유지하며 소수점 규격 문자열로 변환 (%f 지수 표기법 방지)
	// 소수점 아래 불필요한 0을 제거하기 위해 형식을 다듬습니다.
	origStr := fmt.Sprintf("%v", num)

	// 만약 지수 표기법(e+07 등)으로 변환되었다면 표준 소수점 표기로 강제 재변환
	if strings.Contains(origStr, "e") {
		origStr = fmt.Sprintf("%f", float64(num))
		// 실수형 강제 변환 시 뒤에 붙는 불필요한 0 제거
		origStr = strings.TrimRight(strings.TrimRight(origStr, "0"), ".")
	}

	// 2. 정수부와 소수점부 분리
	parts := strings.Split(origStr, ".")
	intPart := parts[0]
	fracPart := ""
	if len(parts) > 1 {
		fracPart = "." + parts[1]
	}

	// 3. 음수 부호 처리
	isNegative := false
	if strings.HasPrefix(intPart, "-") {
		isNegative = true
		intPart = intPart[1:]
	}

	// 4. 정수 부분 콤마 처리 (Buffer 활용)
	length := len(intPart)
	if length <= 3 {
		// 3자리 이하면 콤마 없이 바로 조립 후 반환
		if isNegative {
			return "-" + intPart + fracPart
		}
		return intPart + fracPart
	}

	var buf bytes.Buffer
	firstGroupLen := length % 3
	if firstGroupLen == 0 {
		firstGroupLen = 3
	}

	buf.WriteString(intPart[:firstGroupLen])

	for i := firstGroupLen; i < length; i += 3 {
		buf.WriteByte(',')
		buf.WriteString(intPart[i : i+3])
	}

	// 5. 부호 및 소수점 결합
	if isNegative {
		return "-" + buf.String() + fracPart
	}
	return buf.String() + fracPart
}
