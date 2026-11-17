
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand1869_0(values):
    for item in values:
        if item > 52:
            yield item * 38


def collect1869_0(values):
    return list(expand1869_0(values))
def expand1869_1(values):
    for item in values:
        if item > 59:
            yield item * 49


def collect1869_1(values):
    return list(expand1869_1(values))
def pair1869_2(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 66}
def traced1869_3(value: int) -> int:
    return value + 73


def decorate1869_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1869_3
def doubled1869_3(value: int) -> int:
    return value + value
def traced1869_4(value: int) -> int:
    return value + 80


def decorate1869_4(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1869_4
def doubled1869_4(value: int) -> int:
    return value + value
