package mod

import "fmt"
import "golua/lua"

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

func Register(L *lua.State) {
	L.NewLib("mod", lua.GoFuncs{"test2": test2, "test": test})
}
