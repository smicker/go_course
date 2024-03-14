// NOTE - To run this you must include all source files, like:
// 'go -C "4-PrivatePublic" run main.go samepackage.go'
// or you can just do 'go -C "4-PrivatePublic" run .'
//
// It is the same for building but in the build command you don't
// have to specify the path so you normally just do
// 'go -C "4-PrivatePublic" build' which is the same as
// 'go -C "4-PrivatePublic" build .'
// 'go -C "4-PrivatePublic" build main.go' will not work since
// you would then have to do:
// 'go -C "4-PrivatePublic" build main.go samepackage.go'

package main

import (
	"example/hello/4-PrivatePublic/anotherPackage"
	"fmt"
)

func main() {
	fmt.Println(anotherPackage.PublicVariable)
	// fmt.Println(anotherPackage.privateVariable) // Not reachable

	fmt.Println(anotherPackage.PublicFunction())
	// fmt.Println(anotherPackage.privateFunction()) // Not reachable

	fmt.Println(privateVariableInMain)
	fmt.Println(PublicVariableInMain)

	fmt.Println(privateFuncInMain())
	fmt.Println(PublicFuncInMain())
}
