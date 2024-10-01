// Generated by `wit-bindgen-wrpc-go` 0.9.1. DO NOT EDIT!
package handler

import (
	context "context"
	errors "errors"
	fmt "fmt"
	io "io"
	slog "log/slog"
	utf8 "unicode/utf8"
	wrpc "wrpc.io/go"
)

func Hello(ctx__ context.Context, wrpc__ wrpc.Invoker) (r0__ string, err__ error) {
	var w__ wrpc.IndexWriteCloser
	var r__ wrpc.IndexReadCloser
	w__, r__, err__ = wrpc__.Invoke(ctx__, "wrpc-examples:hello/handler", "hello", nil)
	if err__ != nil {
		err__ = fmt.Errorf("failed to invoke `hello`: %w", err__)
		return
	}
	defer func() {
		if err := r__.Close(); err != nil {
			slog.ErrorContext(ctx__, "failed to close reader", "instance", "wrpc-examples:hello/handler", "name", "hello", "err", err)
		}
	}()
	if cErr__ := w__.Close(); cErr__ != nil {
		slog.DebugContext(ctx__, "failed to close outgoing stream", "instance", "wrpc-examples:hello/handler", "name", "hello", "err", cErr__)
	}
	r0__, err__ = func(r interface {
		io.ByteReader
		io.Reader
	}) (string, error) {
		var x uint32
		var s uint8
		for i := 0; i < 5; i++ {
			slog.Debug("reading string length byte", "i", i)
			b, err := r.ReadByte()
			if err != nil {
				if i > 0 && err == io.EOF {
					err = io.ErrUnexpectedEOF
				}
				return "", fmt.Errorf("failed to read string length byte: %w", err)
			}
			if s == 28 && b > 0x0f {
				return "", errors.New("string length overflows a 32-bit integer")
			}
			if b < 0x80 {
				x = x | uint32(b)<<s
				buf := make([]byte, x)
				slog.Debug("reading string bytes", "len", x)
				_, err = r.Read(buf)
				if err != nil {
					return "", fmt.Errorf("failed to read string bytes: %w", err)
				}
				if !utf8.Valid(buf) {
					return string(buf), errors.New("string is not valid UTF-8")
				}
				return string(buf), nil
			}
			x |= uint32(b&0x7f) << s
			s += 7
		}
		return "", errors.New("string length overflows a 32-bit integer")
	}(r__)
	if err__ != nil {
		err__ = fmt.Errorf("failed to read result 0: %w", err__)
		return
	}
	return
}
