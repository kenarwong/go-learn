package methodsandinterfaces

import (
	"fmt"
	"tour/methodsandinterfaces/interfaces"
	"tour/methodsandinterfaces/methods"
	"tour/methodsandinterfaces/typeassertions"
)

func MethodsAndInterfaces() {
	fmt.Println("Methods and Interfaces")
	fmt.Println("------")
	methods.Methods()
	fmt.Println()
	interfaces.Interfaces()
	fmt.Println()
	typeassertions.TypeAssertions()
	fmt.Println()
}
