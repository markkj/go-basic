package main

import "fmt"

// Open for extension but closed for modification
// Specification

type Color int

const (
	red Color = iota
	green
	blue
)

type Size int

const (
	small Size = iota
	medium
	large
)

type Product struct {
	name  string
	color Color
	size  Size
}

type Filter struct {
}

//if you want to filter by size you have to create new fucntion
func (f *Filter) FilterByColor(products []Product, color Color) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if v.color == color {
			result = append(result, &products[i])
		}
	}
	return result
}

type Specification interface {
	IsSatisfied(p *Product) bool
}

type ColorSpec struct {
	color Color
}

type SizeSpec struct {
	size Size
}

type ColorAndSizeSpec struct {
	first, second Specification
}

func (spec *SizeSpec) IsSatisfied(p *Product) bool {
	return p.size == spec.size
}

func (spec *ColorSpec) IsSatisfied(p *Product) bool {
	return p.color == spec.color
}

func (spec *ColorAndSizeSpec) IsSatisfied(p *Product) bool {
	return spec.first.IsSatisfied(p) && spec.second.IsSatisfied(p)
}

type BetterFilter struct{}

func (f *BetterFilter) Filter(products []Product, spec Specification) []*Product {
	result := make([]*Product, 0)
	for _, v := range products {
		if spec.IsSatisfied(&v) {
			result = append(result, &v)
		}
	}
	return result
}

func main() {
	apple := Product{"Apple", green, small}
	tree := Product{"Tree", green, large}
	house := Product{"House", blue, large}

	products := []Product{apple, tree, house}
	fmt.Println(products)
	filter := Filter{}
	for _, v := range filter.FilterByColor(products, green) {
		fmt.Printf(" - %s is green\n", v.name)
	}

	fmt.Println("-------")

	//Better Filter
	Bfilter := BetterFilter{}
	colorSpec := &ColorSpec{
		color: green,
	}
	for _, v := range Bfilter.Filter(products, colorSpec) {
		fmt.Printf(" - %s is green \n", v.name)
	}

	sizeSpec := &SizeSpec{
		size: small,
	}
	for _, v := range Bfilter.Filter(products, sizeSpec) {
		fmt.Printf(" - %s is small\n", v.name)
	}

	colorAndSizeSpec := &ColorAndSizeSpec{
		first:  colorSpec,
		second: sizeSpec,
	}
	for _, v := range Bfilter.Filter(products, colorAndSizeSpec) {
		fmt.Printf(" - %s size is small and green\n", v.name)
	}
}
