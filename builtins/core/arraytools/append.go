package arraytools

import (
	"github.com/lmorg/murex/lang"
)

func init() {
	lang.GoFunctions["append"] = cmdAppend
	lang.GoFunctions["prepend"] = cmdPrepend
}

func cmdPrepend(p *lang.Process) error {
	dt := p.Stdin.GetDataType()
	p.Stdout.SetDataType(dt)

	if err := p.ErrIfNotAMethod(); err != nil {
		return err
	}

	var array []string

	err := p.Stdin.ReadArray(func(b []byte) {
		array = append(array, string(b))
	})

	if err != nil {
		return err
	}

	array = append(p.Parameters.StringArray(), array...)

	b, err := lang.MarshalData(p, dt, array)
	if err != nil {
		return err
	}

	_, err = p.Stdout.Write(b)
	return err
}

func cmdAppend(p *lang.Process) error {
	dt := p.Stdin.GetDataType()
	p.Stdout.SetDataType(dt)

	if err := p.ErrIfNotAMethod(); err != nil {
		return err
	}

	var array []string

	err := p.Stdin.ReadArray(func(b []byte) {
		array = append(array, string(b))
	})

	if err != nil {
		return err
	}

	array = append(array, p.Parameters.StringArray()...)

	b, err := lang.MarshalData(p, dt, array)
	if err != nil {
		return err
	}

	_, err = p.Stdout.Write(b)
	return err
}
