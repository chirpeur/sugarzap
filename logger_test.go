package sugarzap

import (
	"errors"
	"fmt"
	"testing"
)

func TestInfo(t *testing.T) {
	Info("a b c")
	Infow("haha", "foo", "bar")
	With("foo", "bar").Info("hello")
	With("foo", "bar").Infof("hello %s", "bee")

	logger := WithOptions()
	logger.Info("hi")
	logger.With("foo", "bar").Info("hi")
	logger.WithHash("foo", "bar").Info("hi")
	logger.Infow("hi", "foo", "bar")
	logger.Kind("abc").Info("cvd")
	funC(logger)
	Kind("abc").Info("xyz")
}

func funC(logger Logger) {
	logger.Info("call in func")
}

func TestError(t *testing.T) {
	e1 := errors.New("[it's error 1]")
	e := fmt.Errorf("(warpped error with %w)", e1)
	Error(e)
	Errorf("bad happened %v", e)
	With("foo", "bar").Errorf("bad %v", e)
}
