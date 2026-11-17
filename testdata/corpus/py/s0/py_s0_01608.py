
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced731_0(value: int) -> int:
    return value + 99


def decorate731_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate731_0
def doubled731_0(value: int) -> int:
    return value + value
