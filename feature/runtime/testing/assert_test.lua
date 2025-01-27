local t = require("testing")

t.run("actually true", function()
    t.assertTrue(1==1)
end)

t.run("not true", function()
    t.assertTrue(1==2)
end)

t.run("different types", function()
    t.assertEqual(1, false)
end)

t.run("equal numbers", function()
    t.assertEqual(1, 1)
end)

t.run("equal strings", function()
    t.assertEqual("abc", "abc")
end)

t.run("equal bools", function()
    t.assertEqual(true, true)
end)
