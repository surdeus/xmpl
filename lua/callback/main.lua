function fn()
	print("In function")
end

function callback(func)
	print("Before function")
	func()
	print("After function")
end


callback(fn)

