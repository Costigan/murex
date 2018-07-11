package readline

type EventReturn struct {
	IgnoreKeyPress bool
	ClearHelpers   bool
	CloseReadline  bool
	HintText       []rune
	NewLine        []rune
	NewPos         int
}

// AddEvent registers a new keypress handler
func (rl *Instance) AddEvent(keyPress string, callback func(string, []rune, int) *EventReturn) {
	rl.evtKeyPress[keyPress] = callback
}

// DelEvent deregisters an existing keypress handler
func (rl *Instance) DelEvent(keyPress string) {
	delete(rl.evtKeyPress, keyPress)
}
