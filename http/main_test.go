package main

import (
	"testing"

	"github.com/markkj/go-basic/http/helpers"
)

func TestHelpers(t *testing.T) {
	h := helpers.WEBURL{}
	h.WebURL = "http://google.com"
	h.WebStatus = "200 OK"
	h.AddName("Google")

	if h.Name != "Google" {
		t.Errorf("Expected Google , but got %v", h.Name)
	}
}
