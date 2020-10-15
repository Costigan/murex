package envvars

import (
	"fmt"
	"os"
	"sync"
)

// EnvVars is a table for caching environmental variables to avoid additional
// SYSCALL overhead
type EnvVars struct {
	mutex sync.RWMutex
	vars  []string
}

// New creates a new EnvVar table and populates it the existing environmental
// variables
func New() *EnvVars {
	ev := new(EnvVars)
	ev.vars = os.Environ()
	ev.Unset("PWD")
	shell, err := os.Executable()
	if err != nil {
		shell = os.Args[0]
	}
	ev.Set("SHELL", shell)
	return ev
}

// Get returns an environmental variable
func (ev *EnvVars) Get(key string) string {
	key += "="
	ev.mutex.RLock()
	for i := range ev.vars {
		if len(ev.vars[i]) >= len(key) && ev.vars[i][:len(key)] == key {
			value := ev.vars[i][len(key):]
			ev.mutex.RUnlock()
			return value
		}
	}
	ev.mutex.RUnlock()
	return ""
}

// Set defines an environmental variable
func (ev *EnvVars) Set(key, value string) {
	key += "="
	ev.mutex.Lock()
	for i := range ev.vars {
		if len(ev.vars[i]) >= len(key) && ev.vars[i][:len(key)] == key {
			ev.vars[i] = key + value
			ev.mutex.Unlock()
			return
		}
	}

	ev.mutex.Unlock()
	ev.vars = append(ev.vars, key+value)
}

// Unset a defined environmental variable
func (ev *EnvVars) Unset(key string) error {
	key += "="
	ev.mutex.Lock()
	for i := range ev.vars {
		if len(ev.vars[i]) >= len(key) && ev.vars[i][:len(key)] == key {
			ev.vars[i] = ev.vars[len(ev.vars)-1]
			ev.vars = ev.vars[:len(ev.vars)-1]

			ev.mutex.Unlock()
			return nil
		}
	}

	ev.mutex.Unlock()
	return fmt.Errorf("Cannot unset: Unable to find a key with the name `%s`", key)
}

// Dump returns the entire envvar table
func (ev *EnvVars) Dump() []string {
	ev.mutex.Lock()
	d := make([]string, len(ev.vars))
	copy(d, ev.vars)
	ev.mutex.Unlock()
	return d
}
