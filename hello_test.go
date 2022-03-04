package public_go_module

import (
	"github.com/stretchr/testify/assert"
	"testing"
)
func TestSayHelloFromPublicModule(t *testing.T) {
	assert.Equal(t,SayHelloFromPublicModule("Dominik"), "Hello from public module to Dominik!")
}


func TestSayHelloFromPrivateModule(t *testing.T) {
	assert.Equal(t,SayHelloFromPrivateModule("Dominik"), "Hello from the private go module to Dominik!")
}
