
t = 700

function test(d) 
	for i=500,d,1 do 
		print(i) 
	end 
end

-- lproc.start(test, t)
-- lproc.start(function(t) for i=1,t,1 do print(i) end end, 205)

function send(t) 
	lproc.send("c", 101) 
end

function recv(t) 
	r = lproc.recv("c") 
	print("==========>: ".. r)
end


lproc.start(send)
lproc.start(recv)


lproc.sleep(60)
