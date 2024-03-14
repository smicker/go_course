package anotherPackage

var privateVariable string = "private variable in anotherPackage"
var PublicVariable string = "Public variable in anotherPackage"

func privateFunction() string {
	return privateVariable
}

func PublicFunction() string {
	return privateFunction()
}
