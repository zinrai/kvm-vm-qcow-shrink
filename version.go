package main

import "fmt"

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func printVersion() {
	fmt.Printf("kvm-vm-qcow-shrink %s (commit %s, built %s)\n", version, commit, date)
}
