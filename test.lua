
function test() for i=500,700,1 do print(i) end end

 lproc.start(test)

lproc.start(function() for i=1,200,1 do print(i) end end)


lproc.sleep(60)
