package logx_test

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/mashiike/go-logx"
)

func ExamplePrint() {
	var buf bytes.Buffer
	ctx := logx.WithLogger(context.Background(), log.New(&buf, "context_logger: ", log.Lshortfile))
	log.Print("Hello World!!")
	logx.Print(ctx, strings.Replace("Hello World!?", "o", "", 1))

	fmt.Print(&buf)
	// Output:
	// context_logger: example_test.go:17: Hell World!?
}
