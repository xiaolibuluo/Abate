
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced1276_0(value: int) -> int:
    return value + 6


def decorate1276_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1276_0
def doubled1276_0(value: int) -> int:
    return value + value
def expand1276_1(values):
    for item in values:
        if item > 13:
            yield item * 46


def collect1276_1(values):
    return list(expand1276_1(values))
def traced1276_2(value: int) -> int:
    return value + 20


def decorate1276_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1276_2
def doubled1276_2(value: int) -> int:
    return value + value
