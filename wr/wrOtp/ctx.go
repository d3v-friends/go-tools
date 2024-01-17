package wrOtp

import (
	"context"
	"fmt"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

type Otp struct {
	Issuer string
}

func (x Otp) Generate(accountNm string) (key *otp.Key, err error) {
	return totp.Generate(totp.GenerateOpts{
		Issuer:      x.Issuer,
		AccountName: accountNm,
	})
}

const CtxOtp = "CTX_OTP"

func SetOtp(ctx context.Context, v *Otp) context.Context {
	return context.WithValue(ctx, CtxOtp, v)
}

func GetOtp(ctx context.Context) (res *Otp, err error) {
	var has bool
	if res, has = ctx.Value(CtxOtp).(*Otp); !has {
		err = fmt.Errorf("not found otp in context")
		return
	}
	return
}

func GetOtpP(ctx context.Context) (res *Otp) {
	var err error
	if res, err = GetOtp(ctx); err != nil {
		panic(err)
	}
	return
}
