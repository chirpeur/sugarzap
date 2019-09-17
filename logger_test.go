package sugarzap

import (
	"fmt"
	"testing"
	"time"
)

func TestInfo(t *testing.T) {
	Info("abvc sas s")
	Info("abvc ", 2, 3)
	Info(fmt.Sprintf("abvc,%d,%d ", 2, 3))
	Infof("dfds %d",2)
	Infow(fmt.Sprintf("Failed to fetch URL %d.",3),
		// Structured context as loosely typed key-value pairs.
		"url", "abvc",
		"attempt", 3,
		"backoff", time.Second,
	)

}
