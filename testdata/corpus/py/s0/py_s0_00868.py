
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced394_0(value: int) -> int:
    return value + 83


def decorate394_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate394_0
def doubled394_0(value: int) -> int:
    return value + value
def traced394_1(value: int) -> int:
    return value + 90


def decorate394_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate394_1
def doubled394_1(value: int) -> int:
    return value + value
