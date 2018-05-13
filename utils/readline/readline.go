package readline

import (
	"bytes"
	"errors"
	"os"
	"regexp"
	"unicode/utf8"
)

var rxMultiline *regexp.Regexp = regexp.MustCompile(`[\r\n]+`)

// Readline displays the readline prompt.
// It will return a string (user entered data) or an error.
func (rl *Instance) Readline() (string, error) {
	fd := int(os.Stdin.Fd())
	state, err := MakeRaw(fd)
	if err != nil {
		return "", err
	}
	defer Restore(fd, state)

	print(rl.prompt)

	rl.line = []rune{}
	rl.viUndoHistory = []undoItem{{line: "", pos: 0}}
	rl.pos = 0
	rl.histPos = rl.History.Len()
	rl.modeViMode = vimInsert
	rl.resetHintText()
	rl.resetTabCompletion()

	if len(rl.multisplit) > 0 {
		r := []rune(rl.multisplit[0])
		rl.editorInput(r)
		rl.carridgeReturn()
		if len(rl.multisplit) > 1 {
			rl.multisplit = rl.multisplit[1:]
		} else {
			rl.multisplit = []string{}
		}
		return string(rl.line), nil
	}

	rl.getHintText()
	rl.renderHelpers()

	for {
		rl.viUndoSkipAppend = false
		b := make([]byte, 1024)
		var i int

		if !rl.skipStdinRead {
			var err error
			i, err = os.Stdin.Read(b)
			if err != nil {
				return "", err
			}
		}

		rl.skipStdinRead = false
		r := []rune(string(b))

		if isMultiline(r[:i]) || len(rl.multiline) > 0 {
			rl.multiline = append(rl.multiline, b[:i]...)
			if i == len(b) {
				continue
			}

			if !rl.allowMultiline(rl.multiline) {
				rl.multiline = []byte{}
				continue
			}

			s := string(rl.multiline)
			rl.multisplit = rxMultiline.Split(s, -1)

			r = []rune(rl.multisplit[0])
			rl.modeViMode = vimInsert
			rl.editorInput(r)
			rl.carridgeReturn()
			rl.multiline = []byte{}
			if len(rl.multisplit) > 1 {
				rl.multisplit = rl.multisplit[1:]
			} else {
				rl.multisplit = []string{}
			}
			return string(rl.line), nil
		}

		s := string(r[:i])

		if rl.evtKeyPress[s] != nil {
			ignoreKey, closeReadline, hintText := rl.evtKeyPress[s](s, rl.line, rl.pos)

			if len(hintText) > 0 {
				rl.hintText = hintText
				rl.clearHelpers()
				rl.renderHelpers()
			}
			if ignoreKey {
				continue
			}
			if closeReadline {
				rl.clearHelpers()
				return string(rl.line), nil
			}
		}

		switch b[0] {
		case charCtrlC:
			rl.clearHelpers()
			return "", errors.New(ErrCtrlC)

		case charEOF:
			rl.clearHelpers()
			return "", errors.New(ErrEOF)

		case charTab:
			if rl.modeTabGrid {
				rl.moveTabHighlight(1, 0)
			} else {
				rl.getTabCompletion()
			}

			rl.renderHelpers()
			rl.viUndoSkipAppend = true

		case charCtrlU:
			rl.clearLine()
			rl.resetHelpers()

		case '\r':
			fallthrough
		case '\n':
			if rl.modeTabGrid && len(rl.tcSuggestions) > 0 {
				cell := (rl.tcMaxX * (rl.tcPosY - 1)) + rl.tcPosX - 1
				rl.clearHelpers()
				rl.resetTabCompletion()
				rl.renderHelpers()
				rl.insert([]rune(rl.tcSuggestions[cell]))
				continue
			}
			rl.carridgeReturn()
			return string(rl.line), nil

		case charBackspace, charBackspace2:
			rl.backspace()
			rl.renderHelpers()

		case charEscape:
			rl.escapeSeq(r[:i])

		default:
			rl.editorInput(r[:i])
			if len(rl.multiline) > 0 && rl.modeViMode == vimKeys {
				rl.skipStdinRead = true
			}
		}

		//if !rl.viUndoSkipAppend {
		//	rl.viUndoHistory = append(rl.viUndoHistory, rl.line)
		//}
		rl.undoAppendHistory()
	}
}

func (rl *Instance) escapeSeq(r []rune) {
	switch string(r) {
	case string(charEscape):
		if rl.modeTabGrid {
			rl.clearHelpers()
			rl.resetTabCompletion()
			rl.renderHelpers()
		} else {
			if rl.pos == len(rl.line) && len(rl.line) > 0 {
				rl.pos--
				moveCursorBackwards(1)
			}
			rl.modeViMode = vimKeys
			rl.viIteration = ""
			//rl.viHintVimKeys()
			rl.viHintMessage()
		}
		rl.viUndoSkipAppend = true

	case seqDelete:
		rl.delete()

	case seqUp:
		if rl.modeTabGrid {
			rl.moveTabHighlight(0, -1)
			rl.renderHelpers()
			return
		}
		rl.walkHistory(-1)

	case seqDown:
		if rl.modeTabGrid {
			rl.moveTabHighlight(0, 1)
			rl.renderHelpers()
			return
		}
		rl.walkHistory(1)

	case seqBackwards:
		if rl.modeTabGrid {
			rl.moveTabHighlight(-1, 0)
			rl.renderHelpers()
			return
		}
		if rl.pos > 0 {
			moveCursorBackwards(1)
			rl.pos--
		}
		rl.viUndoSkipAppend = true

	case seqShiftTab:
		if rl.modeTabGrid {
			rl.moveTabHighlight(-1, 0)
			rl.renderHelpers()
			return
		}

	case seqForwards:
		if rl.modeTabGrid {
			rl.moveTabHighlight(1, 0)
			rl.renderHelpers()
			return
		}
		if (rl.modeViMode == vimInsert && rl.pos < len(rl.line)) ||
			(rl.modeViMode != vimInsert && rl.pos < len(rl.line)-1) {
			moveCursorForwards(1)
			rl.pos++
		}
		rl.viUndoSkipAppend = true

	case seqHome, seqHomeSc:
		if rl.modeTabGrid {
			return
		}
		moveCursorBackwards(rl.pos)
		rl.pos = 0
		rl.viUndoSkipAppend = true

	case seqEnd, seqEndSc:
		if rl.modeTabGrid {
			return
		}
		moveCursorForwards(len(rl.line) - rl.pos)
		rl.pos = len(rl.line)
		rl.viUndoSkipAppend = true

	default:
		rl.viUndoSkipAppend = true
	}
}

// editorInput is an unexported function used to determine what mode of text
// entry readline is currently configured for and then update the line entries
// accordingly.
func (rl *Instance) editorInput(r []rune) {
	switch rl.modeViMode {
	case vimKeys:
		rl.vi(r[0])
		rl.viHintMessage()

	case vimDelete:
		rl.vimDelete(r[0])
		rl.viHintMessage()

	case vimReplaceOnce:
		rl.modeViMode = vimKeys
		rl.delete()
		rl.insert([]rune{r[0]})
		rl.viHintMessage()

	case vimReplaceMany:
		for _, char := range r {
			rl.delete()
			rl.insert([]rune{char})
		}
		rl.viHintMessage()

	default:
		rl.insert(r)
	}

	if len(rl.multisplit) == 0 {
		rl.syntaxCompletion()
	}
}

// SetPrompt will define the readline prompt string.
// It also calculates the runes in the string as well as any non-printable escape codes.
func (rl *Instance) SetPrompt(s string) {
	rl.prompt = s

	s = rxAnsiEscSeq.ReplaceAllString(s, "")
	rl.promptLen = utf8.RuneCountInString(s)
}

func (rl *Instance) carridgeReturn() {
	rl.clearHelpers()
	print("\r\n")
	if rl.HistoryAutoWrite {
		var err error
		rl.histPos, err = rl.History.Write(string(rl.line))
		if err != nil {
			print(err.Error() + "\r\n")
		}
	}
}

func isMultiline(r []rune) bool {
	for i := range r {
		if (r[i] == '\r' || r[i] == '\n') && i != len(r)-1 {
			return true
		}
	}
	return false
}

func (rl *Instance) allowMultiline(data []byte) bool {
	printf("\r\nWARNING: %d bytes of multiline data was dumped into the shell!", len(data))
	for {
		print("\r\nDo you wish to proceed (yes|no|preview)? [y/n/p] ")

		b := make([]byte, 1024)

		i, err := os.Stdin.Read(b)
		if err != nil {
			return false
		}

		s := string(b[:i])
		print(s)

		switch s {
		case "y", "Y":
			print("\r\n" + rl.prompt)
			return true

		case "n", "N":
			print("\r\n" + rl.prompt)
			return false

		case "p", "P":
			preview := string(bytes.Replace(data, []byte{'\r'}, []byte{'\r', '\n'}, -1))
			if rl.SyntaxHighlighter != nil {
				preview = rl.SyntaxHighlighter([]rune(preview))
			}
			print("\r\n" + preview)

		default:
			print("\r\nInvalid response. Please answer `y` (yes), `n` (no) or `p` (preview)")
		}
	}
}
