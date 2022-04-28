#!/usr/bin/env python3

from PIL import Image

im = Image.open("../img/cock.jpg")
im.show()

im = im.rotate(90)
im.show()

