package main

var privateVariableInMain string = "private in main"
var PublicVariableInMain string = "Public in main"

func privateFuncInMain() string {
	return privateVariableInMain
}

func PublicFuncInMain() string {
	return privateVariableInMain
}
