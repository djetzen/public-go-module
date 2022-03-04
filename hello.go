package public_go_module

import (
	"fmt"
	private_go_module "github.com/djetzen/private-go-module"
)

func SayHelloFromPrivateModule(name string) string {
	return private_go_module.SayHelloFromPrivateModule(name)
}

func SayHelloFromPublicModule(name string) string {
	return fmt.Sprintf("Hello from public module to %s!",name)
}
