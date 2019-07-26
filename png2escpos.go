package main

import (
	"fmt"
	"os"

	"github.com/mugli/png2escpos/escpos"
)

func main() {
	if len(os.Args) < 2 {
		showHelp()
	}

	imgFile := os.Args[1]
	err := escpos.PrintImage(imgFile, os.Stdout)

	if err != nil {
		fmt.Println("Error printing image: %v", err)
		os.Exit(1)
	}
}

func showHelp() {
	os.Exit(1)
}
