cs = 100
csr = 1

cr = csr
cmr = 100
cdr = 50
cx = 200
cy = 200

function love.update(dt)
	cr = cr	+ cdr*dt
	if(cr >= cmr) then
		cr = csr
	end
end

function love.mousepressed(x, y, button, is, touch)
	cx = x
	cy = y
end

function love.draw()
	love.graphics.circle("fill", cx, cy, cr, cs)
end

