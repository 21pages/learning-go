package log

import (
	"os"
	"testing"
)

func TestSetLevel(t *testing.T) {
	SetLevel(LevelError)
	if infolog.Writer() == os.Stdout || errlog.Writer() != os.Stdout {
		t.Fatal("failed to set log level")
	}
	SetLevel(LevelDisabled)
	if infolog.Writer() == os.Stdout || errlog.Writer() == os.Stdout {
		t.Fatal("failed to set log level")
	}
}
