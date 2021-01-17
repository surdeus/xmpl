-- The array is anything in braces separated by commas.
a = {"Lua", "is", "cool", 1, 2, 3, 10 }
-- Actually in fact they are TABLES, not arrays.

-- Getting the length. (amount of elements that index from 1 to n)
n = #a

for i = 1, n do
	-- table.concat as you can guess out just concatenates all
	-- the elements from table to one string and returns it.
	print(table.concat({"a[", i, "] = ", a[i]}))
end

