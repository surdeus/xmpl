#!/usr/bin/env python3

from tkinter import *

class Window(Frame) :
    def __init__(self, master=None) :
        Frame.__init__(self, master)
        self.master = master
        self.init_window()

    def init_window(self) :
        self.master.title("GUI")
        self.pack(fill=BOTH, expand=1)
        quitButton = Button(self, text="Hello", command=self.hello)
        quitButton.place(x=100, y=100)
    def hello(self):
        print("Hello, World!")

root = Tk()

root.geometry("400x300")

app = Window(root)
root.mainloop()

