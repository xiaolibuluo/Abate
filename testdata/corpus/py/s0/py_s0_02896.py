
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand1328_0(values):
    for item in values:
        if item > 100:
            yield item * 33


def collect1328_0(values):
    return list(expand1328_0(values))
def project1328_1(values):
    select = lambda item: item + 10
    return [select(item) for item in values if item != 0]
def expand1328_2(values):
    for item in values:
        if item > 17:
            yield item * 55


def collect1328_2(values):
    return list(expand1328_2(values))
def traced1328_3(value: int) -> int:
    return value + 24


def decorate1328_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1328_3
def doubled1328_3(value: int) -> int:
    return value + value
