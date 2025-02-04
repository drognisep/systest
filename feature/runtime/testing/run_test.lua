local t = require("testing")

t.run("happy func", function()
	print("Hello!")
end)

t.run("not so happy", function()
	error("Uh-oh!")
end)
