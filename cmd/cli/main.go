package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/YumaFuu/s1m/tui"
)

var Version string

func main() {
	var v, h bool
	flag.BoolVar(&v, "v", false, "Show version")
	flag.BoolVar(&h, "h", false, "Show help")
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "Usage: s1m <flag>")
		fmt.Fprintln(os.Stderr, "Flags:")
		flag.PrintDefaults()
	}

	flag.Parse()

	if v {
		fmt.Println(Version)
		return
	}

	if h {
		flag.Usage()
		return
	}

	tui.Run()
}
