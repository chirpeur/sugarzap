package sugarzap

import (
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
	funC(logger)
}

func funC(logger Logger) {
	logger.Info("call in func")
}
