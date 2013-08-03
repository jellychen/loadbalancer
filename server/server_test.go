package server

import (
	. "github.com/bborbe/assert"
	"testing"
)

func TestImplements(t *testing.T) {
	server, err := NewServer("", []string{"A"})
	if err != nil {
		t.Fatal(err)
	}
	var i *Server
	err = AssertThat(server, Implements(i).Message("check type"))
	if err != nil {
		t.Fatal(err)
	}
}
