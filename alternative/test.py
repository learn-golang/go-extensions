import ctypes

cdll = ctypes.cdll.LoadLibrary("./libadd.so")

print(cdll.add(2, 3))

print(cdll.goroutine("https://python.org", len("https://python.org") + 100 ))
