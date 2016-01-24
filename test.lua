
t = 700

function test(t) for i=500,t,1 do print(i) end end


lproc.start(test, t)

lproc.start(function(t) for i=1,t,1 do print(i) end end, 205)


lproc.sleep(60)
