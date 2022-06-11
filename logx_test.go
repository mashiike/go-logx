package logx_test

import (
	"bytes"
	"context"
	"io"
	"log"
	"strings"
	"testing"

	"github.com/mashiike/go-logx"
)

func TestPrint(t *testing.T) {
	var buf bytes.Buffer
	l := log.New(&buf, "[test]", log.Lshortfile)
	ctx := logx.WithLogger(context.Background(), l)
	logx.Print(ctx, "hoge")
	t.Log(buf.String())
	if !strings.HasPrefix(buf.String(), "[test]logx_test.go:") {
		t.Error("unexpected log prefix")
	}
}

func TestPrintDiscard(t *testing.T) {
	l := log.New(io.Discard, "[test]", log.Lshortfile)
	ctx := logx.WithLogger(context.Background(), l)
	logx.Print(ctx, "hoge")
}

func TestOutput(t *testing.T) {
	var buf bytes.Buffer
	l := log.New(&buf, "[test]", log.Lshortfile)
	log.SetFlags(log.Lshortfile)
	ctx := logx.WithLogger(context.Background(), l)
	testFunc(ctx)
	t.Log(buf.String())
	if !strings.HasPrefix(buf.String(), "[test]logx_test.go:") {
		t.Error("unexpected log prefix")
	}
}
