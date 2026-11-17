
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced390_0(value: int) -> int:
    return value + 31


def decorate390_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate390_0
def doubled390_0(value: int) -> int:
    return value + value
