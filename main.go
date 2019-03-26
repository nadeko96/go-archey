package main

import (
	"log"
	"os"
	"runtime"
	"syscall"
	"time"

	"github.com/nadeko96/go-archey/lib"
)

func bToGb(b uint64) float64 {
	return float64(b) / 1024 / 1024 / 1024
}

func main() {
	var output lib.Output
	output.OS = runtime.GOOS
	output.Username = os.Getenv("USER")
	hostname, err := os.Hostname()
	if err != nil {
		log.Panic(err)
	}
	output.Hostname = hostname

	var info syscall.Sysinfo_t
	err2 := syscall.Sysinfo(&info)
	if err2 != nil {
		log.Panic(err2)
	}
	output.Uptime = time.Duration(info.Uptime) * time.Second
	output.Freeram = bToGb(info.Freeram)
	output.Totalram = bToGb(info.Totalram)
	switch output.OS {
	case "linux":
		lib.PrintLinux(&output)
	case "windows":
		lib.PrintWindows(&output)
	case "darwin":
		lib.PrintOSX(&output)
	default:
		lib.Print(&output)
	}
}
