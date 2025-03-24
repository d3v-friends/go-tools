// Package fnEnv
// 환경변수를 읽어올 때 사용하는 함수 모음
// 환경변수가 없을 때 panic 이 일어나거나 기본값을 출력하는 함수
package fnEnv

import (
	"github.com/d3v-friends/go-tools/fnError"
	"github.com/shopspring/decimal"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"os"
	"strconv"
)

const (
	ErrNotFoundEnv          = "not_found_env"
	ErrInvalidIntValue      = "invalid_int_value"
	ErrInvalidInt64Value    = "invalid_int64_value"
	ErrInvalidDecimalValue  = "invalid_decimal_value"
	ErrInvalidObjectIDValue = "invalid_object_id_value"
)

func String(key string, defaultValues ...string) (value string) {
	value = os.Getenv(key)
	if value == "" {
		if len(defaultValues) == 1 {
			return defaultValues[0]
		}

		panic(fnError.NewFields(ErrNotFoundEnv, map[string]any{
			"key": key,
		}))
	}
	return
}

func Bool(key string, defaultValues ...bool) (res bool) {
	var str = os.Getenv(key)
	if str == "" {
		if len(defaultValues) == 1 {
			return defaultValues[0]
		}
		panic(fnError.NewFields(ErrNotFoundEnv, map[string]any{
			"key": key,
		}))
	}
	return str == "true"
}

func Int64(key string, defaultValues ...int64) int64 {
	var str = os.Getenv(key)
	if str == "" {
		if len(defaultValues) == 1 {
			return defaultValues[0]
		}

		panic(fnError.NewFields(ErrNotFoundEnv, map[string]any{
			"key": key,
		}))
	}

	var d, err = decimal.NewFromString(str)
	if err != nil {
		panic(fnError.NewFields(ErrInvalidInt64Value, map[string]any{
			"key":   key,
			"value": str,
		}))
	}
	return d.IntPart()
}

func Decimal(key string, defaultValues ...decimal.Decimal) decimal.Decimal {
	var str = os.Getenv(key)
	if str == "" {
		if len(defaultValues) == 1 {
			return defaultValues[0]
		}
		panic(fnError.NewFields(ErrNotFoundEnv, map[string]any{
			"key": key,
		}))
	}

	var d, err = decimal.NewFromString(str)
	if err != nil {
		panic(fnError.NewFields(ErrInvalidDecimalValue, map[string]any{
			"key":   key,
			"value": str,
		}))
	}
	return d
}

func ObjectID(key string, defaultValues ...primitive.ObjectID) primitive.ObjectID {
	var str = os.Getenv(key)
	if str == "" {
		if len(defaultValues) == 1 {
			return defaultValues[0]
		}
		panic(fnError.NewFields(ErrNotFoundEnv, map[string]any{
			"key": key,
		}))
	}

	var id, err = primitive.ObjectIDFromHex(str)
	if err != nil {
		panic(fnError.NewFields(ErrInvalidObjectIDValue, map[string]any{
			"key":   key,
			"value": str,
		}))
	}

	return id
}

func Int(key string, defaultValues ...int) int {
	var str = os.Getenv(key)
	if str == "" {
		if len(defaultValues) == 1 {
			return defaultValues[0]
		}
		panic(fnError.NewFields(ErrNotFoundEnv, map[string]any{
			"key": key,
		}))
	}

	var i, err = strconv.Atoi(str)
	if err != nil {
		panic(fnError.NewFields(ErrInvalidIntValue, map[string]any{
			"key":   key,
			"value": str,
		}))
	}
	return i
}
