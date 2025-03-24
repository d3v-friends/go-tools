// Package fnPanic
// 패키지 실행시 if err 체크시 반드시 에러가 나지 않거나, 에러가 날 시 프로그램을 강제로 종료 시키는 함수
// if 문의 중첩을 한줄로 정리 하기 위해서 사용하는 단순한 함수
package fnPanic

import (
	"github.com/d3v-friends/go-tools/fnLogger"
	"github.com/pkg/errors"
)

func On(err error) {
	if err != nil {
		fnLogger.NewLogger().Fatal(err)
		panic(errors.WithStack(err))
	}
}

func Value[T any](value T, err error) T {
	if err != nil {
		fnLogger.NewLogger().Fatal(err)
		panic(errors.WithStack(err))
	}
	return value
}
