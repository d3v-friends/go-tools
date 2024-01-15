package fnMongo

import (
	"context"
	"github.com/d3v-friends/go-tools/fnEnv"
	"github.com/d3v-friends/go-tools/fnPanic"
	"github.com/d3v-friends/go-tools/typ"
	"testing"
)

func TestMigrate(test *testing.T) {
	fnPanic.On(fnEnv.ReadFromFile("../.env"))
	var client = fnPanic.OnValue(ConnectClient(&ConnectClientArgs{
		Host:     fnEnv.Read("MG_HOST"),
		Username: fnEnv.Read("MG_USERNAME"),
		Password: fnEnv.Read("MG_PASSWORD"),
		CodecRegisters: []CodecRegister{
			typ.DecimalRegistry,
		},
	}))

	test.Run("migrate", func(t *testing.T) {
		var ctx = context.TODO()
		ctx = SetCtxDB(ctx, client.Database(fnEnv.Read("MG_DATABASE")))
		var err = RunMigrate(ctx, MangoModel)
		if err != nil {
			t.Fatal(err)
		}
	})
}
