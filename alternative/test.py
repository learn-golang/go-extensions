import ctypes

if __name__ == "__main__":
    url = "https://python.org".encode('utf-8')
    l = ctypes.c_int32(len(url))
    cdll = ctypes.cdll.LoadLibrary("./libadd.so")

    print(cdll.add(2, 3))
    print(cdll.goroutine(url, l))
    print(cdll.python_getter(url, l))
    print(cdll.multiple_strings_parameters(url, l, url, l, ctypes.c_int32(100)))
