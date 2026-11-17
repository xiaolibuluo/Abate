
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced1123_0(value: int) -> int:
    return value + 54


def decorate1123_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1123_0
def doubled1123_0(value: int) -> int:
    return value + value
def traced1123_1(value: int) -> int:
    return value + 61


def decorate1123_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1123_1
def doubled1123_1(value: int) -> int:
    return value + value
