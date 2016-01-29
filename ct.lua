
function foo()
	print("-----------in2")
	coroutine.yield()
	print("-----------hh2")
end



print("------before")
cot.start(function()
		for i = 1, 20, 1 do
			cot.sleep(10)
			print("============= " .. i)
		end
	end)

cot.start(function()
		for i = 1, 20, 1 do
			cot.sleep(15)
			print("----------------- " .. i)
			print("----------------- " .. i)
		end
	end)

print("------after")

-- for i = 1, 30, 1 do
-- 	cot.sleep(30)
-- 	print("------------------------sleep ".. i)
-- end

cot.sleep(5000)