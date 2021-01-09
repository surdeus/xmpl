#include <SDL2/SDL.h>
#include <stdlib.h>

static const char *winTitle = "Hello, World!" ;
static char *argv0 ;

enum{
	WindowX = 100,
	WindowY = 100,
	WindowWidth = 640,
	WindowHeight = 480,
	ExitSuccess = 0,
	ExitError = 1,
};

void
main(int argc, char *argv[])
{
	argv0 = argv[0] ;

	/* Initilization. */
	SDL_Window *win = SDL_CreateWindow(winTitle,
			WindowX, WindowY,
			WindowWidth, WindowHeight,
			SDL_WINDOW_SHOWN) ;
	if(!win){
		fprintf(stderr, "%s: SDL_CreateWindow: %s\n", argv0, SDL_GetError());
		exit(ExitError);
	}

	SDL_Renderer *ren = SDL_CreateRenderer(win, -1,
			SDL_RENDERER_ACCELERATED | SDL_RENDERER_PRESENTVSYNC) ;
	if(!ren){
		fprintf(stderr, "%s: error: SDL_CreateRenderer: %s\n", argv0, SDL_GetError());
		exit(ExitError);
	}

	SDL_Surface *bmp = SDL_LoadBMP("hello.bmp") ;
	if(!bmp){
		fprintf(stderr, "%s: error: SDL_LoadBMP: %s\n", argv0, SDL_GetError());
		exit(ExitError);
	}

	SDL_Texture *tex = SDL_CreateTextureFromSurface(ren, bmp) ;
	SDL_FreeSurface(bmp);
	if(!tex){
		fprintf(stderr, "%s: error: SDL_CreateTextureFromSurface: %s\n", argv0, SDL_GetError());
		exit(ExitError);
	}

	/* Drawing text. */
	SDL_RenderClear(ren);
	SDL_RenderCopy(ren, tex, 0, 0);
	SDL_RenderPresent(ren);

	/* Wait before closing. */
	SDL_Delay(2000);

	/* Cleaning after the window was used. */
	SDL_DestroyTexture(tex);
	SDL_DestroyRenderer(ren);
	SDL_DestroyWindow(win);
	SDL_Quit();

	exit(ExitSuccess);
}

