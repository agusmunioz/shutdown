package shutdown

import (
	"os"
	"os/signal"
)

//Channel defined as global so OnShutdown can be tested.
var interruptChannel = make(chan os.Signal)

//OnShutdown registers a hook shutdown function and hangs until the programs is shutdown by an os interrupt signal.
func OnShutdown(hook func()) {
	signal.Notify(interruptChannel, os.Interrupt)
	<-interruptChannel
	hook()
}
