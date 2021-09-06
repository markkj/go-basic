package main

import "fmt"

type Document struct{}

type Machine interface {
	Print(d Document)
	Fax(d Document)
	Scan(d Document)
}

type SomeOldMachine struct{}

func (old *SomeOldMachine) Print(d Document) {
	// print doc
}

// this method is not necessary
func (old *SomeOldMachine) Fax(d Document) {
	panic("This Machine not support this function")
}

func (old *SomeOldMachine) Scan(d Document) {
	//Scan Doc
}

//IFSP
//Refactor

type Printer interface {
	Print(d Document)
}

type Faxer interface {
	Fax(d Document)
}

type Scaner interface {
	Scan(d Document)
}

type PrinterScanerMachine interface {
	Printer
	Scaner
}

func ShowModel(m PrinterScanerMachine) {
	fmt.Println("model of ", m)
}

type SomeNewMachine struct{}

func (m *SomeNewMachine) Print(d Document) {
	//Print Some Doc
}
func (m *SomeNewMachine) Scan(d Document) {
	//Scan Some Doc
}

func main() {
	newMachine := &SomeNewMachine{}
	ShowModel(newMachine)
}
