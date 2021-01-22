#include <stdlib.h>
#include <stdio.h>

#include <X11/Xlib.h>
#include <X11/Xutil.h>
#include <X11/Xos.h>

enum{
	WinWidth = 200,
	WinHeight = 300,
	ExitKey = 'q',
};

Display *dpy;
int screen;
Window win;
/* Graphical context. */
GC gc;

void
xinit(void)
{
	unsigned long black, white;

	dpy = XOpenDisplay(0) ;
	screen = DefaultScreen(dpy) ;
	black = BlackPixel(dpy, screen) ;
	white = WhitePixel(dpy, screen) ;
	win = XCreateSimpleWindow(dpy, DefaultRootWindow(dpy),
			0, 0,
			WinWidth, WinHeight,
			5,
			black, white) ;
	gc = XCreateGC(dpy, win, 0, 0) ;

	XSetStandardProperties(dpy, win, "Simple window", "Hi.", None, 0, 0, 0);
	XSelectInput(dpy, win, ExposureMask|ButtonPressMask|KeyPressMask);
	XSetBackground(dpy, gc, white);
	XSetForeground(dpy, gc, black);
	XClearWindow(dpy, win);
	XMapRaised(dpy, win);
}

void
xexit(void)
{
	XFreeGC(dpy, gc);
	XDestroyWindow(dpy, win);
	XCloseDisplay(dpy);
}

void
xredraw(void)
{
	XDrawRectangle(dpy, win, gc, 50, 50, 300, 200);
}

void
xwork(void)
{
	XEvent ev;
	KeySym key;
	char buf[BUFSIZ];
	while(1){
		XNextEvent(dpy, &ev);
		if( ev.type==Expose && ev.xexpose.count==0)
			xredraw();
		else if( ev.type==KeyPress
				&& XLookupString( &ev.xkey, buf, sizeof(buf), &key, 0) == 1
			) {
			if(buf[0] == ExitKey) xexit(), exit(0) ;
			printf("'%c'\n", buf[0]);
		} else if (ev.type==ButtonPress)
			printf("%i %i\n", ev.xbutton.x, ev.xbutton.y);
	}
}

int
main(int argc, char *argv[])
{
	xinit();
	xwork();
	xexit();
	return 0 ;
}

