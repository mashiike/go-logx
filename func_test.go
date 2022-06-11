package logx_test

import (
	"context"

	"github.com/mashiike/go-logx"
)

func testFunc(ctx context.Context) {
	logx.Output(ctx, 2, "[info] call from")
}
