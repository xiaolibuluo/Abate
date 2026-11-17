
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand931_0(values):
    for item in values:
        if item > 80:
            yield item * 60


def collect931_0(values):
    return list(expand931_0(values))
def expand931_1(values):
    for item in values:
        if item > 87:
            yield item * 10


def collect931_1(values):
    return list(expand931_1(values))
def traced931_2(value: int) -> int:
    return value + 94


def decorate931_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate931_2
def doubled931_2(value: int) -> int:
    return value + value
