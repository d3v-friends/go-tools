package fnOtp

import (
	"context"
	"fmt"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

const CtxOtp = "CTX_OTP"

type Otp struct {
	Issuer string
}

func (x *Otp) Generate(accountNm string) (key *otp.Key, err error) {
	return totp.Generate(totp.GenerateOpts{
		Issuer:      x.Issuer,
		AccountName: accountNm,
	})
}

func (x *Otp) Verify(otp string, secret string) bool {
	return totp.Validate(otp, secret)
}

func Set(ctx context.Context, v *Otp) context.Context {
	return context.WithValue(ctx, CtxOtp, v)
}

func Get(ctx context.Context) (res *Otp, err error) {
	var has bool
	if res, has = ctx.Value(CtxOtp).(*Otp); !has {
		err = fmt.Errorf("not found otp in context")
		return
	}
	return
}

func GetP(ctx context.Context) (res *Otp) {
	var err error
	if res, err = Get(ctx); err != nil {
		panic(err)
	}
	return
}
