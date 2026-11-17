
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced730_0(value: int) -> int:
    return value + 86


def decorate730_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate730_0
def doubled730_0(value: int) -> int:
    return value + value
