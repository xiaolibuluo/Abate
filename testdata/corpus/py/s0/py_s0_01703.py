
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced779_0(value: int) -> int:
    return value + 44


def decorate779_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate779_0
def doubled779_0(value: int) -> int:
    return value + value
