package main

import (
	"mobin.dev/inteface_practice/consolePrinter"
	"mobin.dev/inteface_practice/pdfPrinter"
)

type printer interface {
	Printer()
}

func main() {
	pdfPrint := pdfPrinter.PDFPrinter{
		Data: "Hello This is PDF printer",
	}

	conPrint := consolePrinter.ConsolePrinter{
		Data: "Hello This is ConsolePrinter",
	}

	printPrinter(pdfPrint)
	printPrinter(conPrint)
}

func printPrinter(p printer) {
	p.Printer()
}
