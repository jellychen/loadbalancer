package connectionhandler

import (
	"testing"
	. "github.com/bborbe/assert"
)

func TestImplements(t *testing.T) {
	connectionhandler := NewConnectionHandler(nil)
	var i *ConnectionHandler
	err := AssertThat(connectionhandler, Implements(i).Message("check type"))
	if err != nil {
		t.Fatal(err)
	}
}
