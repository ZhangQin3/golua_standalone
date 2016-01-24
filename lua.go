package main

import (
	"github.com/ZhangQin3/golua/lua"
	// "lua/lproc"
	"lua/model"
	"os"
)

func main() {
	L := lua.NewState()

	mod.Register(L)
	// lproc.Register(L)
	lua.RegisterLProc(L)
	lua.MainGo(L, os.Args)
}
