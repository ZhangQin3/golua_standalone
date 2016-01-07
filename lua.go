package main

import "os"
import "golua/lua"
import "golua_standalone/model"

func main() {
	L := lua.NewState()
	// Register go functions test and test2 in ./model/mod.go file to lua
	mod.Register(L)
	lua.MainGo(L, os.Args)
}
