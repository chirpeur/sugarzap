package sugarzap

import (
	"testing"
)

func TestInfo(t *testing.T) {
	Info("a b c")
	Infow("haha", "foo", "bar")
	With("foo", "bar").Info("hello")
	With("foo", "bar").Infof("hello %s", "bee")

	logger := WithOptions(AddCallerSkip(1))
	logger.Info("hi")
	logger.With("foo", "bar").Info("hi")
	logger.WithHash("foo", "bar").Info("hi")
	logger.Infow("hi", "foo", "bar")
}
