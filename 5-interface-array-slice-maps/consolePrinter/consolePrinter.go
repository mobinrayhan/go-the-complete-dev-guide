package consolePrinter

import (
	"fmt"
	"time"
)

type ConsolePrinter struct {
	Data     string
	CreateAt time.Time
}

func (p ConsolePrinter) Printer() {
	fmt.Println(p.Data)
}
