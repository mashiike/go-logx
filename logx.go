package logx

import (
	"context"
	"fmt"
	"io"
	"os"
)

// SetOutput sets the output destination for the standard logger tied to the context.
// For more details, https://pkg.go.dev/log#SetOutput
func SetOutput(ctx context.Context, w io.Writer) {
	Default(ctx).SetOutput(w)
}

// Flags returns the output flags for the standard logger tied to the context.
// For more details, https://pkg.go.dev/log#Flags
func Flags(ctx context.Context) int {
	return Default(ctx).Flags()
}

// SetFlags sets the output flags for the standard logger tied to the context.
// For more details, https://pkg.go.dev/log#SetFlags
func SetFlags(ctx context.Context, flag int) {
	Default(ctx).SetFlags(flag)
}

// Prefix returns the output prefix for the standard logger tied to the context.
// For more details, https://pkg.go.dev/log#Prefix
func Prefix(ctx context.Context) string {
	return Default(ctx).Prefix()
}

// SetPrefix sets the output prefix for the standard logger tied to the context.
// For more details, https://pkg.go.dev/log#SetPrefix
func SetPrefix(ctx context.Context, prefix string) {
	Default(ctx).SetPrefix(prefix)
}

// Writer returns the output destination for the standard logger tied to the context.
// For more details, https://pkg.go.dev/log#Writer
func Writer(ctx context.Context) io.Writer {
	return Default(ctx).Writer()
}

// Print calls Output to print to the standard logger tied to the context.
// For more details, https://pkg.go.dev/log#Print
func Print(ctx context.Context, v ...any) {
	Default(ctx).Output(2, fmt.Sprint(v...))
}

// Printf calls Output to print to the standard logger tied to the context.
// For more details, https://pkg.go.dev/log#Printf
func Printf(ctx context.Context, format string, v ...any) {
	Default(ctx).Output(2, fmt.Sprintf(format, v...))
}

// Println calls Output to print to the standard logger tied to the context.
// For more details, https://pkg.go.dev/log#Println
func Println(ctx context.Context, v ...any) {
	Default(ctx).Output(2, fmt.Sprintln(v...))
}

// Fatal is equivalent to Print() followed by a call to os.Exit(1)
// For more details, https://pkg.go.dev/log#Fatal
func Fatal(ctx context.Context, v ...any) {
	Default(ctx).Output(2, fmt.Sprint(v...))
	os.Exit(1)
}

// Fatalf is equivalent to Printf() followed by a call to os.Exit(1).
func Fatalf(ctx context.Context, format string, v ...any) {
	Default(ctx).Output(2, fmt.Sprintf(format, v...))
	os.Exit(1)
}

// Fatalln is equivalent to Println() followed by a call to os.Exit(1).
func Fatalln(ctx context.Context, v ...any) {
	Default(ctx).Output(2, fmt.Sprintln(v...))
	os.Exit(1)
}

// Panic is equivalent to Print() followed by a call to panic().
func Panic(ctx context.Context, v ...any) {
	s := fmt.Sprint(v...)
	Default(ctx).Output(2, s)
	panic(s)
}

// Panicf is equivalent to Printf() followed by a call to panic().
func Panicf(ctx context.Context, format string, v ...any) {
	s := fmt.Sprintf(format, v...)
	Default(ctx).Output(2, s)
	panic(s)
}

// Panicln is equivalent to Println() followed by a call to panic().
func Panicln(ctx context.Context, v ...any) {
	s := fmt.Sprintln(v...)
	Default(ctx).Output(2, s)
	panic(s)
}

// Output writes the output for a logging event with logger tied to the context.
func Output(ctx context.Context, calldepth int, s string) error {
	return Default(ctx).Output(calldepth+1, s) // +1 for this frame.
}
