package main

import "fmt"

type Image interface {
	Draw()
}

type Bitmap struct {
	filename string
}

func NewBitMap(filename string) *Bitmap {
	fmt.Println("Loading image from ", filename)
	return &Bitmap{filename: filename}
}

func (b *Bitmap) Draw() {
	fmt.Println("Drawing Image", b.filename)
}

func DrawImage(image Image) {
	fmt.Println("About to draw the image")
	image.Draw()
	fmt.Println("Done drawing the image")

}

type LazyBitmap struct {
	filename string
	bitmap   *Bitmap
}

func NewLazyBitmap(filename string) *LazyBitmap {
	return &LazyBitmap{filename: filename}
}

func (b *LazyBitmap) Draw() {
	if b.bitmap == nil {
		b.bitmap = NewBitMap(b.filename)
	}
	b.bitmap.Draw()
}

func main() {
	bmp := NewBitMap("mark.png")
	DrawImage(bmp)

	bmps := NewLazyBitmap("mark.png")
	DrawImage(bmps)
}
