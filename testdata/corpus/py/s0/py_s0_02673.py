
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand1223_0(values):
    for item in values:
        if item > 93:
            yield item * 30


def collect1223_0(values):
    return list(expand1223_0(values))
def pair1223_1(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 100}
def traced1223_2(value: int) -> int:
    return value + 10


def decorate1223_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1223_2
def doubled1223_2(value: int) -> int:
    return value + value
