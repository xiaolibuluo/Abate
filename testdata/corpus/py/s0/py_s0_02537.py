
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced1169_0(value: int) -> int:
    return value + 70


def decorate1169_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1169_0
def doubled1169_0(value: int) -> int:
    return value + value
