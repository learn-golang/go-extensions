import ctypes

if __name__ == "__main__":
    cdll = ctypes.cdll.LoadLibrary("./libadd.so")

    print(cdll.add(2, 3))

    print(cdll.goroutine("https://python.org".encode('utf-8'), ctypes.c_int32(len("https://python.org"))))
