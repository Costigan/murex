package main

import (
	"fmt"
	"testing"

	"github.com/lmorg/murex/lang"
)

func BenchmarkA(b *testing.B) {
	lang.InitEnv()

	block := fmt.Sprintf(`a [1..%d] -> foreach i { out "iteration $i of %d" }`, b.N, b.N)

	_, err := lang.ShellProcess.Fork(lang.F_NO_STDIN | lang.F_NO_STDOUT | lang.F_NO_STDERR).Execute([]rune(block))
	if err != nil {
		b.Error(err.Error())
	}
}

func BenchmarkJa(b *testing.B) {
	lang.InitEnv()

	block := fmt.Sprintf(`ja [1..%d] -> foreach i { out "iteration $i of %d" }`, b.N, b.N)

	_, err := lang.ShellProcess.Fork(lang.F_NO_STDIN | lang.F_NO_STDOUT | lang.F_NO_STDERR).Execute([]rune(block))
	if err != nil {
		b.Error(err.Error())
	}
}

func BenchmarkCSV(b *testing.B) {
	lang.InitEnv()

	block := []rune(`tout csv "murex,foo,bar\n1,2,3\na,b,c\nz,y,x\n" -> [ :foo ]`)

	for i := 0; i < b.N; i++ {
		_, err := lang.ShellProcess.Fork(lang.F_NO_STDIN | lang.F_NO_STDOUT | lang.F_NO_STDERR).Execute(block)
		if err != nil {
			b.Error(err.Error())
		}
	}
}
