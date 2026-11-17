
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand846_0(values):
    for item in values:
        if item > 42:
            yield item * 14


def collect846_0(values):
    return list(expand846_0(values))
def pair846_1(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 49}
def traced846_2(value: int) -> int:
    return value + 56


def decorate846_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate846_2
def doubled846_2(value: int) -> int:
    return value + value
def pair846_3(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 63}
def pair846_4(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 70}
