package wrJwt

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
	"time"
)

func TestJWT(test *testing.T) {
	var secret, issuer = "secret1234!1234", "dev_friends"

	test.Run("test", func(t *testing.T) {
		var v1 = NewJwt(secret, issuer)
		var data = testData{
			Id:       primitive.NewObjectID(),
			Username: "dev_friends",
			Subject:  "sign_in",
		}

		var jwtToken, err = v1.Encode(data)
		if err != nil {
			t.Fatal(err)
		}

		fmt.Printf("jwt_token=%s\n", jwtToken)

		var id string
		if id, err = v1.Decode(jwtToken); err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, data.Id.Hex(), id)
	})

	test.Run("expire", func(t *testing.T) {
		var v1 = NewJwt(secret, issuer, time.Second*2)

		var data = testData{
			Id:       primitive.NewObjectID(),
			Username: "dev_friends",
			Subject:  "sign_in",
		}

		var jwtToken, err = v1.Encode(data)
		if err != nil {
			t.Fatal(err)
		}

		fmt.Printf("jwt_token=%s\n", jwtToken)
		if _, err = v1.Decode(jwtToken); err != nil {
			t.Fatal(err)
		}

		time.Sleep(time.Second * 4)

		_, err = v1.Decode(jwtToken)
		fmt.Printf("err=%s\n", err.Error())
		if err == nil {
			t.Fatal("not expired jwt token")
		}

	})

}

type testData struct {
	Id       primitive.ObjectID
	Username string
	Subject  string
}

func (x testData) GetID() string {
	return x.Id.Hex()
}

func (x testData) GetAudience() []string {
	return []string{
		x.Username,
	}
}

func (x testData) GetSubject() string {
	return x.Subject
}
