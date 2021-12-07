#!/bin/env python3

import sys
import requests
from bs4 import BeautifulSoup as bs

if len(sys.argv) != 2 :
	exit(1)

url = sys.argv[1]

soup = bs(requests.get(url).text, "lxml")
for tag in soup.find_all("a", href=True) :
    print(tag.attrs["href"])

