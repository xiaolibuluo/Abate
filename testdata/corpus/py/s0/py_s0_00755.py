
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def pair345_0(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 28}
def traced345_1(value: int) -> int:
    return value + 35


def decorate345_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate345_1
def doubled345_1(value: int) -> int:
    return value + value
