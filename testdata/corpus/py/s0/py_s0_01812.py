
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def pair829_0(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 15}
def traced829_1(value: int) -> int:
    return value + 22


def decorate829_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate829_1
def doubled829_1(value: int) -> int:
    return value + value
