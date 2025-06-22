package main

import "fmt"

// クライアント目線で利用したいインターフェース
type Printer interface {
	Print()
}

type OldPrinter struct{}

func (op *OldPrinter) OldPrint() { // Printerインターフェースを満たしていない
	fmt.Println("古いプリンターで印刷しています")
}

type PrinterAdapter struct {
	Old *OldPrinter
}

func (pa *PrinterAdapter) Print() {
	fmt.Println("アダプターを使って変換中...")
	pa.Old.OldPrint()
}

func UsePrinter(p Printer) {
	p.Print()
}

func main() {
	old := &OldPrinter{}
	adapter := &PrinterAdapter{Old: old}
	UsePrinter(adapter)
}
