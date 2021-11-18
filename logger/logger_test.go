package logger_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/anschwa/giftopotamus/logger"
)

var b bytes.Buffer

func init() {
	logger.Init(&b)
}

func testLevel(t *testing.T, want, got string) {
	if !strings.Contains(got, want) {
		t.Errorf("want: %s; got: %s", want, got)
	}
}

func TestDebug(t *testing.T) {
	b.Reset()

	logger.Debug("abc")
	testLevel(t, "DEBU abc", b.String())
}

func Test_Debugf(t *testing.T) {
	b.Reset()

	logger.Debugf("%s %d", "abc", 123)
	testLevel(t, "DEBU abc 123", b.String())
}

func Test_Info(t *testing.T) {
	b.Reset()

	logger.Info("abc")
	testLevel(t, "INFO abc", b.String())
}

func Test_Infof(t *testing.T) {
	b.Reset()

	logger.Infof("%s %d", "abc", 123)
	testLevel(t, "INFO abc 123", b.String())
}

func Test_Warn(t *testing.T) {
	b.Reset()

	logger.Warn("abc")
	testLevel(t, "WARN abc", b.String())
}

func Test_Warnf(t *testing.T) {
	b.Reset()

	logger.Warnf("%s %d", "abc", 123)
	testLevel(t, "WARN abc 123", b.String())
}

func Test_Error(t *testing.T) {
	b.Reset()

	logger.Error("abc")
	testLevel(t, "ERRO abc", b.String())
}

func TestErrorf(t *testing.T) {
	b.Reset()

	logger.Errorf("%s %d", "abc", 123)
	testLevel(t, "ERRO abc 123", b.String())
}

// Cannot test log.Fatal

func Test_Panic(t *testing.T) {
	b.Reset()

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("no panic")
		}
	}()

	logger.Panic("abc")
	testLevel(t, "PANI abc", b.String())
}

func Test_Panicf(t *testing.T) {
	b.Reset()

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("no panic")
		}
	}()

	logger.Panicf("%s %d", "abc", 123)
	testLevel(t, "PANI abc 123", b.String())
}
