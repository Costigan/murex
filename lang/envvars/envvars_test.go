package envvars_test

import (
	"testing"

	"github.com/lmorg/murex/lang/envvars"
)

var envvarTestName = "MUREX_TEST_TestEnvVars_FOOBAR_"

func TestEnvVars(t *testing.T) {
	ev := envvars.New()
	v := ev.Get(envvarTestName)
	if v != "" {
		t.Errorf("'%s' != '' ('%s')", envvarTestName, v)
	}

	ev.Set(envvarTestName, "foobar")
	v = ev.Get(envvarTestName)
	if v != "foobar" {
		t.Errorf("'%s' != 'foobar' ('%s')", envvarTestName, v)
	}

	err := ev.Unset(envvarTestName)
	if err != nil {
		t.Errorf("%s", err.Error())
	}
	v = ev.Get(envvarTestName)
	if v != "" {
		t.Errorf("'%s' != '' ('%s')", envvarTestName, v)
	}
}
