#!/usr/bin/env python3
import sys
import time
import keyboard

DELAY=5
time.sleep(DELAY)

f = sys.stdin
try:
    l = f.readline()
    while l:
        keyboard.write(l)
        l = f.readline()
finally:
    print("Done")
sys.exit(0)
