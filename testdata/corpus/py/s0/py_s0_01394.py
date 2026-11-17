
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced635_0(value: int) -> int:
    return value + 15


def decorate635_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate635_0
def doubled635_0(value: int) -> int:
    return value + value
def traced635_1(value: int) -> int:
    return value + 22


def decorate635_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate635_1
def doubled635_1(value: int) -> int:
    return value + value
