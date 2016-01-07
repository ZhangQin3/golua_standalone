package main

import "golua/lua"
import "os"

func main() {
	L := lua.NewState()
	lua.MainGo(L, os.Args)
}
