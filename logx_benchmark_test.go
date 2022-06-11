package logx_test

import (
	"context"
	"io"
	"log"
	"testing"

	"github.com/mashiike/go-logx"
)

func BenchmarkLogDefaultLogger(b *testing.B) {
	log.SetOutput(io.MultiWriter(io.Discard, io.Discard))
	log.SetPrefix("[prefix]")
	log.SetFlags(log.Ldate | log.LUTC)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		log.Println("[info] hoge")
	}
}

func BenchmarkLogxDefaultLogger(b *testing.B) {
	log.SetOutput(io.MultiWriter(io.Discard, io.Discard))
	log.SetPrefix("[prefix]")
	log.SetFlags(log.Ldate | log.LUTC)
	ctx := context.Background()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logx.Println(ctx, "[info] hoge")
	}
}

func BenchmarkLogxContextLogger(b *testing.B) {
	l := log.New(io.MultiWriter(io.Discard, io.Discard), "[prefix]", log.Ldate|log.LUTC)
	ctx := logx.WithLogger(context.Background(), l)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logx.Println(ctx, "[info] hoge")
	}
}
