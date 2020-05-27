package shutdown

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

//Test the shutdown hook when an os.Interrupt signal is notified.
func TestShutdown(t *testing.T) {

	waitShutdown := make(chan bool)

	hook := func() {
		waitShutdown <- true
	}

	go OnShutdown(hook)

	interruptChannel <- os.Interrupt

	select {
	case s := <-waitShutdown:
		assertion := assert.New(t)
		assertion.True(s, "Unexpected value")
	case <-time.After(3 * time.Second):
		t.Fatalf("Timeout waiting for shutdown hook to be executed")
	}

}
