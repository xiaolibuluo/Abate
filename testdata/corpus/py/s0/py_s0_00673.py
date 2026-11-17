
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand312_0(values):
    for item in values:
        if item > 84:
            yield item * 58


def collect312_0(values):
    return list(expand312_0(values))
def pair312_1(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 91}
def pair312_2(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 98}
def traced312_3(value: int) -> int:
    return value + 8


def decorate312_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate312_3
def doubled312_3(value: int) -> int:
    return value + value
def pair312_4(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 15}
def traced312_5(value: int) -> int:
    return value + 22


def decorate312_5(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate312_5
def doubled312_5(value: int) -> int:
    return value + value
