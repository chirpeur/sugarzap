package sugarzap

import (
	"fmt"
	"testing"
)

func TestInfo(t *testing.T) {
	With("foo","bar").Infof("hello %s", "bee")
	Info("a b c")
	Info("a", "b", "c")
	_globalLogger.Info("a b c")
	_globalLogger.Info("a", "b", "c")

	//Error("abvc ", 2, 3)
	//Info(fmt.Sprintf("abvc,%d,%d ", 2, 3))
	//Infof("dfds %d", 2)
	//Infow(fmt.Sprintf("Failed to fetch URL %d.", 3),
	//	// Structured context as loosely typed key-value pairs.
	//	"url", "abvc",
	//	"attempt", 3,
	//	"backoff", time.Second,
	//)
	msg := fmt.Sprint("a", "b")
	fmt.Println(msg)
	fmt.Println(fmt.Sprint([]interface{}{msg}))
	//err := fmt.Errorf("this is err: %w", errors.New("foo"))
	//Errorf("this is err: %w", errors.New("foo"))
	Fatal(msg)

}
