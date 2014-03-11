package server

import (
	"testing"
	. "github.com/bborbe/assert"
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
