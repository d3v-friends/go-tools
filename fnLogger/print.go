package fnLogger

import (
	"context"
	"fmt"
	"log"
	"runtime"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func printLog(
	ctx context.Context,
	stdout *log.Logger,
	level LogLevel,
	message any,
	showLineNumber bool,
	groups ...LogGroup,
) {
	var id, err = GetID(ctx)
	if err != nil {
		id = &CtxID{
			Id:        primitive.NilObjectID.Hex(),
			CreatedAt: time.Now(),
		}
	}

	var group = NilLogGroup
	if len(groups) == 1 {
		group = groups[0]
	}

	var loc, line = getLocation()
	var color = ColorKeyGray

	var str = fmt.Sprintf("[%s]%s[%s]\t%s",
		level.log(),
		group.String(),
		id.Id,
		color.log(stringify(message)),
	)

	if showLineNumber {
		str = fmt.Sprintf("%s [%s](%d)", str, loc, line)
	}

	stdout.Print(str)
}

func getLocation() (loc string, line int) {
	var pc uintptr
	pc, loc, line, _ = runtime.Caller(4)

	var fnName = runtime.FuncForPC(pc).Name()
	var lastSlash = strings.LastIndex(fnName, "/")
	if lastSlash < 0 {
		lastSlash = 0
	}

	var lastDot = strings.LastIndexByte(fnName[lastSlash:], '.') + lastSlash
	loc = fnName[:lastDot]
	return
}
