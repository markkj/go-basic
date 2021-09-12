package main

import (
	"fmt"
	"strings"
)

type ExpressionVisitor interface {
	VisitDoubleExpression(de *DoubleExpression)
	VisitAdditionExpression(ae *AdditionExpression)
}

type Expression interface {
	Accept(ev ExpressionVisitor)
}

type DoubleExpression struct {
	value float64
}

func (d *DoubleExpression) Accept(ev ExpressionVisitor) {
	ev.VisitDoubleExpression(d)
}

type AdditionExpression struct {
	left, right Expression
}

func (a *AdditionExpression) Accept(ev ExpressionVisitor) {
	ev.VisitAdditionExpression(a)
}

type ExpressionPrinter struct {
	sb strings.Builder
}

func (e *ExpressionPrinter) VisitDoubleExpression(de *DoubleExpression) {
	e.sb.WriteString(fmt.Sprintf("%g", de.value))
}

func (e *ExpressionPrinter) VisitAdditionExpression(ae *AdditionExpression) {
	e.sb.WriteString("(")
	ae.left.Accept(e)
	e.sb.WriteString("+")
	ae.right.Accept(e)
	e.sb.WriteString(")")
}

func NewExpressionPrinter() *ExpressionPrinter {
	return &ExpressionPrinter{strings.Builder{}}
}

func (e *ExpressionPrinter) String() string {
	return e.sb.String()
}

type ExpressionEvaluator struct {
	result float64
}

func (e *ExpressionEvaluator) VisitDoubleExpression(de *DoubleExpression) {
	e.result = de.value
}

func (e *ExpressionEvaluator) VisitAdditionExpression(ae *AdditionExpression) {
	ae.left.Accept(e)
	x := e.result
	ae.right.Accept(e)
	x += e.result
	e.result = x
}

func main() {
	// 1+(2+3)
	e := &AdditionExpression{
		&DoubleExpression{1},
		&AdditionExpression{
			left:  &DoubleExpression{2},
			right: &DoubleExpression{3},
		},
	}
	ep := NewExpressionPrinter()
	e.Accept(ep)
	// ep.VisitAdditionExpression(e)
	// fmt.Println(ep.String())

	ee := &ExpressionEvaluator{}
	e.Accept(ee)

	fmt.Printf("%s = %g\n", ep, ee.result)

}
