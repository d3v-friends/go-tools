package grCouponPoint

import (
	"database/sql/driver"
	"fmt"
)

type RequestStatus string

const (
	RequestStatusRequested    RequestStatus = "REQUESTED"     // 요청됨
	RequestStatusInProcessing RequestStatus = "IN_PROCESSING" // 수정중
	RequestStatusFail         RequestStatus = "FAIL"          // 실패
	RequestStatusSucceed      RequestStatus = "SUCCEED"       // 성공
)

var RequestStatusAll = []RequestStatus{
	RequestStatusRequested,
	RequestStatusInProcessing,
	RequestStatusFail,
	RequestStatusSucceed,
}

func (x *RequestStatus) String() string {
	return string(*x)
}

func (x *RequestStatus) Valid() bool {
	for _, status := range RequestStatusAll {
		if *x == status {
			return true
		}
	}
	return false
}

func (x *RequestStatus) Scan(src any) (err error) {
	switch v := src.(type) {
	case string:
		*x = RequestStatus(v)
		return
	case []byte:
		*x = RequestStatus(v)
		return
	default:
		err = fmt.Errorf("invalid src type")
		return
	}
}

func (x *RequestStatus) Value() (res driver.Value, err error) {
	res = string(*x)
	return
}
