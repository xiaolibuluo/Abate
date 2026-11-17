
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced1472_0(value: int) -> int:
    return value + 32


def decorate1472_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1472_0
def doubled1472_0(value: int) -> int:
    return value + value
def traced1472_1(value: int) -> int:
    return value + 39


def decorate1472_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1472_1
def doubled1472_1(value: int) -> int:
    return value + value
def expand1472_2(values):
    for item in values:
        if item > 46:
            yield item * 26


def collect1472_2(values):
    return list(expand1472_2(values))
