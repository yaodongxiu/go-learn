package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"strings"
	"testing"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randomString(n int) string {
	b := make([]byte, n)
	for i := 0; i < n; i++ {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func benchmark(b *testing.B, f func(int, string) string) {
	str := randomString(10)
	for i := 0; i < b.N; i++ {
		f(10000, str)
	}
}

func plusConcat(n int, str string) string {
	s := ""
	for i := 0; i < n; i++ {
		s += str
	}
	return s
}

func sprintfConcat(n int, str string) string {
	s := ""
	for i := 0; i < n; i++ {
		s = fmt.Sprintf("%s%s", s, str)
	}
	return s
}
func builderConcat(n int, str string) string {
	var builder strings.Builder
	for i := 0; i < n; i++ {
		builder.WriteString(str)
	}
	return builder.String()
}

func preBuilderConcat(n int, str string) string {
	var builder strings.Builder
	builder.Grow(n * len(str))
	for i := 0; i < n; i++ {
		builder.WriteString(str)
	}
	return builder.String()
}

func bufferConcat(n int, str string) string {
	var buffer bytes.Buffer
	for i := 0; i < n; i++ {
		buffer.WriteString(str)
	}
	return buffer.String()
}

func preBufferConcat(n int, str string) string {
	var buffer bytes.Buffer
	buffer.Grow(n * len(str))
	for i := 0; i < n; i++ {
		buffer.WriteString(str)
	}
	return buffer.String()
}
func byteConcat(n int, str string) string {
	b := make([]byte, 0)
	for i := 0; i < n; i++ {
		b = append(b, str...)
	}
	return string(b)
}
func preByteConcat(n int, str string) string {
	b := make([]byte, 0, len(str)*n)
	for i := 0; i < n; i++ {
		b = append(b, str...)
	}
	return string(b)
}

func BenchmarkPlusConcat(b *testing.B)       { benchmark(b, plusConcat) }
func BenchmarkSprintfConcat(b *testing.B)    { benchmark(b, sprintfConcat) }
func BenchmarkBuilderConcat(b *testing.B)    { benchmark(b, builderConcat) }
func BenchmarkPreBuilderConcat(b *testing.B) { benchmark(b, preBuilderConcat) }
func BenchmarkBufferConcat(b *testing.B)     { benchmark(b, bufferConcat) }
func BenchmarkPreBufferConcat(b *testing.B)  { benchmark(b, preBufferConcat) }
func BenchmarkByteConcat(b *testing.B)       { benchmark(b, byteConcat) }
func BenchmarkPreByteConcat(b *testing.B)    { benchmark(b, preByteConcat) }
