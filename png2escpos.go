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

	if os.Args[1] == "help" || os.Args[1] == "--help" || os.Args[1] == "-h" {
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
	help := `


	▄▄▄▄▄▄▄ ▄▄  ▄ ▄▄▄▄▄▄▄ 
	█ ▄▄▄ █ ▄▀▄ █ █ ▄▄▄ █ 
	█ ███ █ █▄▄▀  █ ███ █ 
	█▄▄▄▄▄█ ▄ ▄ ▄ █▄▄▄▄▄█ 
	▄▄▄▄  ▄ ▄▀█  ▄  ▄▄▄ ▄ 
	▀▄█ ▀▀▄▀ ▀▀ █ ▄▄▀█▄▄█ 
	███▄▄▄▄▄█ █▄▀ ▄█ ▀ ▀  
	▄▄▄▄▄▄▄ ▀▄▄  ▀▄▄█▄▄▀  
	█ ▄▄▄ █  ▄▀█▀█  ▀ █ ▀ 
	█ ███ █ █▄█▄ █▄▀   ▀  
	█▄▄▄▄▄█ █▄▄██▄▀█▄ ▄ ▀ 



------
Usage:

Quickly converts PNG files to ESC/POS format and write to stdout, 
	for printing on Epson thermal point-of-sale printers.

This utility removes transparency from image, makes it grayscale
and then encodes in ESC/POS format.

./png2escpos <filename.png>:
	Binary output in ESC/POS format will be written directly to stdout.

	You can pipe this output directly into an Epson compatible thermal printer with:
		Linux:	./png2escpos <file.png> > /dev/usb/lp0
		macOS:	./png2escpos <file.png> | lpr -P NAME_OF_PRINTER

		Or, if you have a network printer listening at 192.168.1.100:9100 
			you can use socat to forward from stdin to network like this:
			
			Linux:	./png2escpos <file.png> > socat STDIN TCP4:192.168.1.100:9100
			macOS:	./png2escpos <file.png> | socat STDIN TCP4:192.168.1.100:9100
			
			(you may have to install socat on macOS with brew install socat first)
		
Other commands:

	help, --help, -h:	
		Shows this message
`
	fmt.Println(help)
	os.Exit(0)
}
