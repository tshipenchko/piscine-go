#!/usr/bin/env python3
import os

text = """
dr-------x 1986-01-05 00:00 0
-r------w- 1986-11-13 00:01 1
-rw----r-- 1988-03-05 00:10 2
lrwxrwxrwx 1990-02-16 00:11 3 -> 0
-r-x--x--- 1990-10-07 01:00 4
-r--rw---- 1990-11-07 01:01 5
-r--rw---- 1991-02-08 01:10 6
-r-x--x--- 1991-03-08 01:11 7
-rw----r-- 1994-05-20 10:00 8
-r------w- 1994-06-10 10:01 9
dr-------x 1995-04-10 10:10 A
"""
for line in text.split("\n"):
    if not line:
        continue

    mod, date, time, name = line.split(" ", 3)

    f_date = date.replace("-", "")
    f_time = time.replace(":", "")
    command = f"TZ=utc lrwxrwxrwx 2023-02-21 08touch -t {f_date}{f_time} src/{name}"
    print(command)
    os.system(command)
