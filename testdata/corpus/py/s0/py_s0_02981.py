
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand1365_0(values):
    for item in values:
        if item > 96:
            yield item * 48


def collect1365_0(values):
    return list(expand1365_0(values))
def traced1365_1(value: int) -> int:
    return value + 6


def decorate1365_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1365_1
def doubled1365_1(value: int) -> int:
    return value + value
