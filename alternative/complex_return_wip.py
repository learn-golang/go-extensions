import ctypes


class complex_return_return(ctypes.Structure):
    _fields_ = [
        ("status", ctypes.POINTER(ctypes.c_char_p)),
        ("code", ctypes.POINTER(ctypes.c_char_p)),
        # ("another_code", ctypes.POINTER(ctypes.c_char_p)),
    ]


if __name__ == "__main__":
    url = "https://python.org".encode('utf-8')
    l = ctypes.c_int32(len(url))

    cdll = ctypes.cdll.LoadLibrary("./libadd.so")

    pointer_to_complex = ctypes.POINTER(complex_return_return)
    cdll.complex_return.restype = pointer_to_complex
    complex_return = cdll.complex_return(url, l)
    print(str(complex_return.contents.status, 'utf-8').split('object'))
