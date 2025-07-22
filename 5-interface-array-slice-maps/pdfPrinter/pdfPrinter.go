package pdfPrinter

import (
	"fmt"
	"time"
)

type PDFPrinter struct {
	Data     string
	CreateAt time.Time
}

func (p PDFPrinter) Printer() {
	fmt.Println(p.Data)
}
