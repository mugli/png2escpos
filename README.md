# png2escpos

Quickly convert PNG files to ESC/POS format, for printing on Epson thermal point-of-sale printers.

This utility removes transparency from image, makes it grayscale and then encodes in ESC/POS format.

Written in pure Go/Golang.

---

## Installation

Download binary from [release tab](https://github.com/mugli/png2escpos/releases).

---

## Usage

```
./png2escpos <filename.png>:
	Binary output in ESC/POS format will be written directly to stdout.

	You can pipe this output directly into an Epson compatible thermal printer with:
		Linux:	./png2escpos <file.png> > /dev/usb/lp0
		macOS:	./png2escpos <file.png> | lpr -P NAME_OF_PRINTER

		Or, if you have a network printer listening at 192.168.1.100:9100
			you can use socat to forward from stdin to network like this:

			Linux:	./png2escpos <file.png> | socat STDIN TCP4:192.168.1.100:9100
			macOS:	./png2escpos <file.png> | socat STDIN TCP4:192.168.1.100:9100

			(you may have to install socat on macOS with brew install socat first)

Other commands:

	help, --help, -h:
		Shows this message
```

---

## Acknowledgement

[NielsLeenheer/EscPosEncoder](https://github.com/NielsLeenheer/EscPosEncoder) - For the rasterization and ESC/POS encoding in JavaScript.

[twg/png2escpos](https://github.com/twg/png2escpos) - For CLI usage pattern.

---

## License

MIT
