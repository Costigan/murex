package string

import (
	"bufio"
	"fmt"
	"strconv"

	"github.com/lmorg/murex/lang"
)

/*func index(p *lang.Process, params []string) error {
	match := make(map[string]bool)
	for i := range params {
		match[params[i]] = true
	}

	err := p.Stdin.ReadMap(p.Config, func(key, value string, last bool) {
		if p.IsNot {
			if !match[key] {
				p.Stdout.Writeln([]byte(value))
			}
		} else {
			if match[key] {
				p.Stdout.Writeln([]byte(value))
			}
		}
	})

	return err
}*/

func index(p *lang.Process, params []string) error {
	lines := make(map[int]bool)
	for i := range params {
		num, err := strconv.Atoi(params[i])
		if err != nil {
			return fmt.Errorf("Parameter, `%s`, isn't an integer. %s", params[i], err)
		}
		lines[num] = true
	}

	var (
		i   int
		err error
	)

	scanner := bufio.NewScanner(p.Stdin)
	for scanner.Scan() {
		if lines[i] != p.IsNot {
			_, err = p.Stdout.Writeln(scanner.Bytes())
			if err != nil {
				break
			}
		}
		i++
	}

	return err
}
