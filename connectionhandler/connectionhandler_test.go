package connectionhandler

import (
	. "github.com/bborbe/assert"
	"testing"
)

func TestImplements(t *testing.T) {
	connectionhandler := NewConnectionHandler(nil)
	var i *ConnectionHandler
	err := AssertThat(connectionhandler, Implements(i).Message("check type"))
	if err != nil {
		t.Fatal(err)
	}
}
