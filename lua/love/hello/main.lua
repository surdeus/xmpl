w = 640
h = 480
x = 4
xd = 1
y = 5
yd = 2

function love.keypressed(key, scancode, isrepeat)
	love.event.quit()
end

function love.draw()
	love.graphics.print("Hello, World!", x, y)
	if (x<=0 or x>=w) then
		xd = -xd
	end
	if (y<=0 or y>=h) then
		yd = -yd
	end
	x = x + xd
	y = y + yd
end

