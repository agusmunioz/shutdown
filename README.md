# Shutdown

A Go module for executing a behaviour when application is shutdown by an external signal.

```go
package main

import (
     "fmt"
     "github.com/agusmunioz/shutdown"
)

func main() {

    shutdownHook := func() {
        fmt.Println("Graceful shutdown")
    }
 
    shutdown.OnShutdown(shutdownHook)
 	
}
```
