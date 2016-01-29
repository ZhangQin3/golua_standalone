package main

import (
	// "fmt"
	"github.com/ZhangQin3/golua/lua"
	// "lua_standalone/model"
	"os"
)

func main() {
	L := lua.NewState()

	// mod.Register(L)
	// lua.RegisterLProc(L)
	lua.RegisterCothread(L)
	lua.MainGo(L, os.Args)
}
