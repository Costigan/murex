package management

import (
	"io/ioutil"
	"os"
	"time"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/proc/runmode"
	"github.com/lmorg/murex/lang/ref"
)

func init() {
	lang.GoFunctions["source"] = cmdSource
}

func cmdSource(p *lang.Process) error {
	var (
		block []rune
		name  string
		err   error
		b     []byte
	)

	if p.IsMethod {
		b, err = p.Stdin.ReadAll()
		if err != nil {
			return err
		}
		block = []rune(string(b))
		name = "<stdin>"

	} else {
		block, err = p.Parameters.Block(0)
		if err == nil {
			b = []byte(string(block))
			name = "N/A"

		} else {
			// get block from file
			name, err = p.Parameters.String(0)
			if err != nil {
				return err
			}

			file, err := os.Open(name)
			if err != nil {
				return err
			}

			b, err = ioutil.ReadAll(file)
			if err != nil {
				return err
			}
			block = []rune(string(b))
		}
	}

	module := quickHash(name + time.Now().String())
	fileRef := &ref.File{Source: ref.History.AddSource(name, "source/"+module, b)}

	p.RunMode = runmode.Normal
	fork := p.Fork(lang.F_FUNCTION | lang.F_NEW_MODULE | lang.F_NO_STDIN)

	fork.Name = p.Name
	fork.FileRef = fileRef
	p.ExitNum, err = fork.Execute(block)
	return err
}
