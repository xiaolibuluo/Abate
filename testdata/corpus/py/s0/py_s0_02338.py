
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced1075_0(value: int) -> int:
    return value + 12


def decorate1075_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1075_0
def doubled1075_0(value: int) -> int:
    return value + value
def pair1075_1(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 19}
