package main

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"
)

var entryCount = 0

type Journal struct {
	entries []string
}

func (j *Journal) toString() string {
	return strings.Join(j.entries, ",")
}

func (j *Journal) AddEntry(text string) int {
	entryCount++
	entry := fmt.Sprintf("%d: %s",
		entryCount, text)

	j.entries = append(j.entries, entry)
	return entryCount
}

func (j *Journal) RemoveEntry(index int) {

}

//bad
func (j *Journal) Save(filename string) {
	_ = ioutil.WriteFile(filename,
		[]byte(j.toString()), 0664)
}

//bad
func (j *Journal) Load(filename string) {

}

//bad
func (j *Journal) LoadFromWeb(url *url.URL) {

}

// Should seperate responsibility to persistance
type Persistance struct {
	LineSep string
}

func (p *Persistance) SaveToFile(filename string, j *Journal) {
	_ = ioutil.WriteFile(filename,
		[]byte(strings.Join(j.entries, p.LineSep)), 0664)
}

func main() {
	j := Journal{}
	j.AddEntry("DATA01")
	j.AddEntry("DATA02")
	fmt.Println(j.toString())
	p := Persistance{
		LineSep: "-",
	}
	p.SaveToFile("./testSave01.txt", &j)

}
