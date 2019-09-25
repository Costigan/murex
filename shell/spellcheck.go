package shell

import (
	"strings"

	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/ansi"
)

func spellCheck(line []rune) []rune {
	r := line
	enabled, err := lang.ShellProcess.Config.Get("shell", "spellcheck-enabled", types.Boolean)
	if err != nil || !enabled.(bool) {
		return r
	}

	block, err := lang.ShellProcess.Config.Get("shell", "spellcheck-block", types.CodeBlock)
	if err != nil || len(block.(string)) == 0 {
		return r
	}

	fork := lang.ShellProcess.Fork(lang.F_FUNCTION | lang.F_CREATE_STDIN | lang.F_CREATE_STDOUT | lang.F_CREATE_STDERR)
	fork.Name = "(spellcheck"
	fork.Stdin.SetDataType(types.Generic)
	_, err = fork.Stdin.Writeln([]byte(string(r)))
	if err != nil && debug.Enabled {
		lang.ShellProcess.Stderr.Writeln([]byte(err.Error()))
		return r
	}

	_, err = fork.Execute([]rune(block.(string)))
	if err != nil && debug.Enabled {
		lang.ShellProcess.Stderr.Writeln([]byte(err.Error()))
		return r
	}

	b, err := fork.Stderr.ReadAll()
	if err != nil && debug.Enabled {
		lang.ShellProcess.Stderr.Writeln([]byte(err.Error()))
	}
	if len(b) != 0 && debug.Enabled {
		lang.ShellProcess.Stderr.Writeln([]byte(err.Error()))
	}

	err = fork.Stdout.ReadArray(func(bword []byte) {
		sword := string(bword)
		r = []rune(strings.ReplaceAll(string(r), sword, ansi.ExpandConsts("{UNDERLINE}"+sword+"{UNDEROFF}")))
	})
	if err != nil && debug.Enabled {
		lang.ShellProcess.Stderr.Writeln([]byte(err.Error()))
	}

	return r
}
