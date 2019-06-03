//////////////////////////////////////////////////////////////////////
//
// Given is a mock process which runs indefinitely and blocks the
// program. Right now the only way to stop the program is to send a
// SIGINT (Ctrl-C). Killing a process like that is not graceful, so we
// want to try to gracefully stop the process first.
//
// Change the program to do the following:
//   1. On SIGINT try to gracefully stop the process using
//          `proc.Stop()`
//   2. If SIGINT is called again, just kill the program (last resort)
//

package main

import (
	"os"
	"os/signal"
	"syscall"

	signint "github.com/kennykarnama/go-concurrency-exercises/4-graceful-sigint/signit"
)

func main() {
	proc := signint.MockProcess{}

	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	var counter int
	go func() {
		for {
			select {
			case <-sig:
				counter++
				//fmt.Printf("Interrupted\n")
				go func() {
					proc.Stop()
				}()
				if counter == 2 {
					os.Exit(1)
				}

			}
		}
		// os.Exit(1)

	}()

	proc.Run()
}
