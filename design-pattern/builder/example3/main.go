package main

import (
	"fmt"
	"strings"
)

type Email struct {
	from, to, subject, body string
}

type EmailBuilder struct {
	email Email
}

func (b *EmailBuilder) From(from string) *EmailBuilder {
	if !strings.Contains(from, "@") {
		panic("Email is incorrect")
	}
	b.email.from = from
	return b
}

func (b *EmailBuilder) To(to string) *EmailBuilder {
	if !strings.Contains(to, "@") {
		panic("Email is incorrect")
	}
	b.email.to = to
	return b
}

func (b *EmailBuilder) Subject(subject string) *EmailBuilder {
	b.email.subject = subject
	return b
}

func (b *EmailBuilder) Body(body string) *EmailBuilder {
	b.email.body = body
	return b
}

type build func(*EmailBuilder)

func SendEmail(action build) {
	builder := EmailBuilder{}
	action(&builder)
	fmt.Print("Sent :", builder)
}

func main() {
	SendEmail(func(eb *EmailBuilder) {
		eb.From("mark@gmail.com").To("Mark2@gmail.com").
			Subject("Hello This my Email").Body("Hello Again")
	})
}
