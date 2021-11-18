package log

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	l, ok := New(Info).(*logger)

	assert.True(t, ok)
	assert.NotNil(t, l)
	assert.NotNil(t, l.logger)
}

func TestLogger_ChangeVerbosity(t *testing.T) {
	l := new(logger)
	l.ChangeVerbosity(Info)

	assert.Equal(t, Info, l.verbosity)
}

func TestLogger_Debug(t *testing.T) {
	l := New(Debug)
	l.Debug("value", 27)
}

func TestLogger_Debugf(t *testing.T) {
	l := New(Debug)
	l.Debugf("Hello, %s!", "World")
}

func TestLogger_Info(t *testing.T) {
	l := New(Info)
	l.Info("value", 27)
}

func TestLogger_Infof(t *testing.T) {
	l := New(Info)
	l.Infof("Hello, %s!", "World")
}

func TestLogger_Warn(t *testing.T) {
	l := New(Warn)
	l.Warn("value", 27)
}

func TestLogger_Warnf(t *testing.T) {
	l := New(Warn)
	l.Warnf("Hello, %s!", "World")
}

func TestLogger_Error(t *testing.T) {
	l := New(Error)
	l.Error("value", 27)
}

func TestLogger_Errorf(t *testing.T) {
	l := New(Error)
	l.Errorf("Hello, %s!", "World")
}

func TestLogger_Fatal(t *testing.T) {
	v := Verbosity(99)
	l := New(v)
	l.Fatal("value", 27)
}

func TestLogger_Fatalf(t *testing.T) {
	v := Verbosity(99)
	l := New(v)
	l.Fatalf("Hello, %s!", "World")
}
