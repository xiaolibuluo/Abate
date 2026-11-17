
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced537_0(value: int) -> int:
    return value + 99


def decorate537_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate537_0
def doubled537_0(value: int) -> int:
    return value + value
