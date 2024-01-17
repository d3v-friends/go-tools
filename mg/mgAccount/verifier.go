package mgAccount

import (
	"context"
	"fmt"
	"github.com/d3v-friends/go-tools/wr/wrOtp"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

type VerifierMode string

func (x VerifierMode) String() string {
	return string(x)
}

func (x VerifierMode) Valid() bool {
	for _, mode := range VerifierModeAll {
		if x == mode {
			return true
		}
	}
	return false
}

const (
	VerifierModeCompare VerifierMode = "compare"
	VerifierModeOtp     VerifierMode = "otp"
)

var VerifierModeAll = []VerifierMode{
	VerifierModeCompare,
	VerifierModeOtp,
}

/*------------------------------------------------------------------------------------------------*/

type Verifier map[VerifierKey]*VerifierValue

/*------------------------------------------------------------------------------------------------*/

type VerifierKey string

func (x VerifierKey) String() string {
	return string(x)
}

/*------------------------------------------------------------------------------------------------*/

type VerifierValue struct {
	Key   string
	Value string
	Mode  VerifierMode
}

func (x VerifierValue) Verify(answer string) bool {
	switch x.Mode {
	case VerifierModeCompare:
		return x.Value == answer
	case VerifierModeOtp:
		return totp.Validate(answer, x.Value)
	default:
		panic(fmt.Errorf("invalid verifier mode: mode=%s", x.Mode))
	}
}

// NewVerifierValueOTP
// ctx 에 *wrOtp.Otp 가 있어야 한다.
func NewVerifierValueOTP(ctx context.Context, accountNm string) (res *VerifierValue, err error) {
	var generator = wrOtp.GetOtpP(ctx)

	var key *otp.Key
	if key, err = generator.Generate(accountNm); err != nil {
		return
	}

	res = &VerifierValue{
		Key:   key.URL(),
		Value: key.Secret(),
		Mode:  VerifierModeOtp,
	}
	return
}

/*------------------------------------------------------------------------------------------------*/
