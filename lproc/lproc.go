package lproc

import (
	"github.com/ZhangQin3/golua/lua"
	"time"
)

func start(L *lua.State) int {
	L.PStart()
	return 0
}

func send(L *lua.State) int {
	return L.Send()
}

func recv(L *lua.State) int {
	return L.Recv()
}

func sleep(L *lua.State) int {
	L.GCheckFunctionArgs("sleep", 1)
	n := time.Duration(L.CheckInteger(1))
	time.Sleep(n * time.Millisecond)
	return 0
}

func Register(L *lua.State) {
	L.NewLib("lproc", lua.GoFuncs{"start": start, "send": send, "recv": recv, "sleep": sleep})
}
