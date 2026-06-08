package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/metacubex/mihomo/constant"
	"github.com/metacubex/mihomo/hub"
	"github.com/metacubex/mihomo/hub/executor"
	"github.com/metacubex/mihomo/log"
)

var (
	homeDir    string
	configFile string
)

func init() {
	flag.StringVar(&homeDir, "d", "", "set configuration directory")
	flag.StringVar(&configFile, "f", "", "specify configuration file")
	flag.Parse()
}

func main() {
	if homeDir != "" {
		constant.SetHomeDir(homeDir)
	}

	if configFile != "" {
		constant.SetConfig(configFile)
	}

	cfg, err := executor.ParseWithPath(constant.Path.Config())
	if err != nil {
		log.Errorln("parse config error: %s", err.Error())
		os.Exit(1)
	}

	hub.ApplyConfig(cfg)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh

	executor.Shutdown()
}
