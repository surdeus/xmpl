#include <raylib.h>

int
main(void)
{
	int x=0, y=0;
	const int w =  800, h = 450 ;
	InitWindow(w, h, "Hello, World!");
	while(!WindowShouldClose()){
	BeginDrawing();
		ClearBackground(RAYWHITE);
		DrawText("Hello, Raylib!", x%w, y%h, 20, BLACK);
		++x; ++y;
	EndDrawing();
	}
	return 0 ;
}
