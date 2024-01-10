package fnJWT

import (
	"fmt"
	"github.com/d3v-friends/go-pure/fnParams"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type JWT[DATA IfJwtData] struct {
	secret []byte
	issuer string
	expire time.Duration
}

type IfJwtData interface {
	GetID() string
}

func NewJWT[DATA IfJwtData](
	secret, issuer string,
	expires ...time.Duration,
) *JWT[DATA] {
	var expire = fnParams.Get(expires)
	if expire == 0 {
		expire = -1
	}

	return &JWT[DATA]{
		secret: []byte(secret),
		issuer: issuer,
		expire: expire,
	}
}

func (x *JWT[DATA]) Encode(data DATA) (res string, err error) {
	var now = time.Now()
	var nowNumericDate = &jwt.NumericDate{
		Time: now,
	}

	var claims = &jwt.RegisteredClaims{
		Issuer:    x.issuer,
		NotBefore: nowNumericDate,
		IssuedAt:  nowNumericDate,
		ID:        data.GetID(),
	}

	if 0 < x.expire {
		claims.ExpiresAt = &jwt.NumericDate{
			Time: now.Add(x.expire),
		}
	}

	var token = jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	if res, err = token.SignedString(x.secret); err != nil {
		return
	}

	return
}

func (x *JWT[DATA]) Decode(str string) (res primitive.ObjectID, err error) {
	var claims = new(jwt.RegisteredClaims)
	var token *jwt.Token
	if token, err = jwt.ParseWithClaims(str, claims, func(token *jwt.Token) (interface{}, error) {
		return x.secret, nil
	}); err != nil {
		return
	}

	if !token.Valid {
		claims = nil
		err = fmt.Errorf("invalid jwt_str: str=%s", str)
		return
	}

	if res, err = primitive.ObjectIDFromHex(claims.ID); err != nil {
		return
	}

	return
}
