package mod

import "fmt"
import "github.com/ZhangQin3/golua/lua"
import "unsafe"

func test(L *lua.State) int {
	fmt.Println("hello world! from go!")
	return 0
}

func test2(L *lua.State) int {
	arg := L.CheckInteger(-1)
	argfrombottom := L.CheckInteger(1)
	fmt.Print("test2 arg: ")
	fmt.Println(arg)
	fmt.Print("from bottom: ")
	fmt.Println(argfrombottom)
	return 0
}

type Userdata struct {
	a, b int
}

func new(L *lua.State) int {
	rawptr := L.NewUserdata(uintptr(unsafe.Sizeof(Userdata{})))
	L.LSetMetaTable("mod.metatable")

	var ptr *Userdata
	ptr = (*Userdata)(rawptr)
	ptr.a = 2
	ptr.b = 3

	fmt.Println(ptr)
	return 1
}

func set(L *lua.State) int {
	ptr := (*Userdata)(L.ToUserdata(-1))
	ptr.a = 99
	ptr.b = 88

	return 0
}

func get(L *lua.State) int {
	ptr := (*Userdata)(L.ToUserdata(-1))

	L.PushInteger(int64(ptr.a))
	L.PushInteger(int64(ptr.b))

	return 2
}

func Register(L *lua.State) {
	L.GNewMetaTable("mod.metatable")
	L.SetFuncs(lua.GoFuncs{"set": set, "get": get})

	L.NewLib("mod", lua.GoFuncs{"test": test, "test2": test2, "new": new})
}
