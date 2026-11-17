
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced1467_0(value: int) -> int:
    return value + 64


def decorate1467_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1467_0
def doubled1467_0(value: int) -> int:
    return value + value
def expand1467_1(values):
    for item in values:
        if item > 71:
            yield item * 41


def collect1467_1(values):
    return list(expand1467_1(values))
def expand1467_2(values):
    for item in values:
        if item > 78:
            yield item * 52


def collect1467_2(values):
    return list(expand1467_2(values))
