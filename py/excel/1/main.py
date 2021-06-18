#!/bin/env python3

from openpyxl import Workbook

file = "file.xlsx"

wb = Workbook()
sheet = wb.active

sheet["A1"] = "Hello"
sheet["B2"] = "World"
wb.save(filename=file)
