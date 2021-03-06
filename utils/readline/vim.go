package readline

import (
	"strconv"
	"strings"
)

type viMode int

const (
	vimInsert viMode = iota
	vimReplaceOnce
	vimReplaceMany
	vimDelete
	vimKeys
)

func (rl *Instance) vi(r rune) {
	switch r {
	case 'a':
		if len(rl.line) > 0 {
			moveCursorForwards(1)
			rl.pos++
		}
		rl.modeViMode = vimInsert
		rl.viIteration = ""
		rl.viUndoSkipAppend = true

	case 'A':
		if len(rl.line) > 0 {
			moveCursorForwards(len(rl.line) - rl.pos)
			rl.pos = len(rl.line)
		}
		rl.modeViMode = vimInsert
		rl.viIteration = ""
		rl.viUndoSkipAppend = true

	case 'b':
		rl.viUndoSkipAppend = true
		vii := rl.getViIterations()
		for i := 1; i <= vii; i++ {
			rl.moveCursorByAdjust(rl.viJumpB(tokeniseLine))
		}

	case 'B':
		rl.viUndoSkipAppend = true
		vii := rl.getViIterations()
		for i := 1; i <= vii; i++ {
			rl.moveCursorByAdjust(rl.viJumpB(tokeniseSplitSpaces))
		}

	case 'd':
		rl.modeViMode = vimDelete
		rl.viUndoSkipAppend = true

	case 'D':
		moveCursorBackwards(rl.pos)
		print(strings.Repeat(" ", len(rl.line)))

		moveCursorBackwards(len(rl.line) - rl.pos)
		rl.line = rl.line[:rl.pos]
		rl.echo()

		moveCursorBackwards(2)
		rl.pos--
		rl.viIteration = ""

	case 'e':
		rl.viUndoSkipAppend = true
		vii := rl.getViIterations()
		for i := 1; i <= vii; i++ {
			rl.moveCursorByAdjust(rl.viJumpE(tokeniseLine))
		}

	case 'E':
		rl.viUndoSkipAppend = true
		vii := rl.getViIterations()
		for i := 1; i <= vii; i++ {
			rl.moveCursorByAdjust(rl.viJumpE(tokeniseSplitSpaces))
		}

	case 'h':
		if rl.pos > 0 {
			moveCursorBackwards(1)
			rl.pos--
		}
		rl.viUndoSkipAppend = true

	case 'i':
		rl.modeViMode = vimInsert
		rl.viIteration = ""
		rl.viUndoSkipAppend = true

	case 'I':
		rl.modeViMode = vimInsert
		rl.viIteration = ""
		rl.viUndoSkipAppend = true
		moveCursorBackwards(rl.pos)
		rl.pos = 0

	case 'l':
		if (rl.modeViMode == vimInsert && rl.pos < len(rl.line)) ||
			(rl.modeViMode != vimInsert && rl.pos < len(rl.line)-1) {
			moveCursorForwards(1)
			rl.pos++
		}
		rl.viUndoSkipAppend = true

	case 'p':
		// paste after
		rl.viUndoSkipAppend = true
		rl.pos++
		moveCursorForwards(1)
		vii := rl.getViIterations()
		for i := 1; i <= vii; i++ {
			rl.insert([]rune(rl.viYankBuffer))
		}
		rl.pos--
		moveCursorBackwards(1)

	case 'P':
		// paste before
		rl.viUndoSkipAppend = true
		vii := rl.getViIterations()
		for i := 1; i <= vii; i++ {
			rl.insert([]rune(rl.viYankBuffer))
		}

	case 'r':
		rl.modeViMode = vimReplaceOnce
		rl.viIteration = ""
		rl.viUndoSkipAppend = true

	case 'R':
		rl.modeViMode = vimReplaceMany
		rl.viIteration = ""
		rl.viUndoSkipAppend = true

	case 'u':
		rl.undoLast()
		rl.viUndoSkipAppend = true

	case 'v':
		rl.clearHelpers()
		var multiline []rune
		if rl.GetMultiLine == nil {
			multiline = rl.line
		} else {
			multiline = rl.GetMultiLine(rl.line)
		}

		new, err := rl.launchEditor(multiline)
		if err != nil || len(new) == 0 || string(new) == string(multiline) {
			rl.viUndoSkipAppend = true
			return
		}
		rl.clearLine()
		rl.multiline = []byte(string(new))

	case 'w':
		rl.viUndoSkipAppend = true
		vii := rl.getViIterations()
		for i := 1; i <= vii; i++ {
			rl.moveCursorByAdjust(rl.viJumpW(tokeniseLine))
		}

	case 'W':
		rl.viUndoSkipAppend = true
		vii := rl.getViIterations()
		for i := 1; i <= vii; i++ {
			rl.moveCursorByAdjust(rl.viJumpW(tokeniseSplitSpaces))
		}

	case 'x':
		vii := rl.getViIterations()
		for i := 1; i <= vii; i++ {
			rl.delete()
		}
		if rl.pos == len(rl.line) && len(rl.line) > 0 {
			moveCursorBackwards(1)
			rl.pos--
		}

	case 'y', 'Y':
		rl.viYankBuffer = string(rl.line)
		rl.viUndoSkipAppend = true
		//rl.hintText = []rune("-- LINE YANKED --")
		//rl.renderHelpers()

	case '[':
		rl.viUndoSkipAppend = true
		rl.moveCursorByAdjust(rl.viJumpPreviousBrace())

	case ']':
		rl.viUndoSkipAppend = true
		rl.moveCursorByAdjust(rl.viJumpNextBrace())

	case '$':
		moveCursorForwards(len(rl.line) - rl.pos)
		rl.pos = len(rl.line)
		rl.viUndoSkipAppend = true

	case '%':
		rl.viUndoSkipAppend = true
		rl.moveCursorByAdjust(rl.viJumpBracket())

	default:
		if r <= '9' && '0' <= r {
			rl.viIteration += string(r)
		}
		rl.viUndoSkipAppend = true

	}
}

func (rl *Instance) getViIterations() int {
	i, _ := strconv.Atoi(rl.viIteration)
	if i < 1 {
		i = 1
	}
	rl.viIteration = ""
	return i
}

func (rl *Instance) viHintMessage() {
	switch rl.modeViMode {
	case vimKeys:
		rl.hintText = []rune("-- VIM KEYS -- (press `i` to return to normal editing mode)")
	case vimInsert:
		rl.hintText = []rune("-- INSERT --")
	case vimReplaceOnce:
		rl.hintText = []rune("-- REPLACE CHARACTER --")
	case vimReplaceMany:
		rl.hintText = []rune("-- REPLACE --")
	case vimDelete:
		rl.hintText = []rune("-- DELETE --")
	default:
		rl.getHintText()
	}

	rl.clearHelpers()
	rl.renderHelpers()
}

func (rl *Instance) viJumpB(tokeniser func([]rune, int) ([]string, int, int)) (adjust int) {
	split, index, pos := tokeniser(rl.line, rl.pos)
	switch {
	case len(split) == 0:
		return
	case index == 0 && pos == 0:
		return
	case pos == 0:
		adjust = len(split[index-1])
	default:
		adjust = pos
	}
	return adjust * -1
}

func (rl *Instance) viJumpE(tokeniser func([]rune, int) ([]string, int, int)) (adjust int) {
	split, index, pos := tokeniser(rl.line, rl.pos)
	if len(split) == 0 {
		return
	}

	word := rTrimWhiteSpace(split[index])

	switch {
	case len(split) == 0:
		return
	case index == len(split)-1 && pos >= len(word)-1:
		return
	case pos >= len(word)-1:
		word = rTrimWhiteSpace(split[index+1])
		adjust = len(split[index]) - pos
		adjust += len(word) - 1
	default:
		adjust = len(word) - pos - 1
	}
	return
}

func (rl *Instance) viJumpW(tokeniser func([]rune, int) ([]string, int, int)) (adjust int) {
	split, index, pos := tokeniser(rl.line, rl.pos)
	switch {
	case len(split) == 0:
		return
	case index+1 == len(split):
		adjust = len(rl.line) - 1 - rl.pos
	default:
		adjust = len(split[index]) - pos
	}
	return
}

func (rl *Instance) viJumpPreviousBrace() (adjust int) {
	if rl.pos == 0 {
		return 0
	}

	for i := rl.pos - 1; i != 0; i-- {
		if rl.line[i] == '{' {
			return i - rl.pos
		}
	}

	return 0
}

func (rl *Instance) viJumpNextBrace() (adjust int) {
	if rl.pos >= len(rl.line)-1 {
		return 0
	}

	for i := rl.pos + 1; i < len(rl.line); i++ {
		if rl.line[i] == '{' {
			return i - rl.pos
		}
	}

	return 0
}

func (rl *Instance) viJumpBracket() (adjust int) {
	split, index, pos := tokeniseBrackets(rl.line, rl.pos)
	switch {
	case len(split) == 0:
		return
	case pos == 0:
		adjust = len(split[index])
	default:
		adjust = pos * -1
	}
	return
}
