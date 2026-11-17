
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand401_0(values):
    for item in values:
        if item > 77:
            yield item * 10


def collect401_0(values):
    return list(expand401_0(values))
def traced401_1(value: int) -> int:
    return value + 84


def decorate401_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate401_1
def doubled401_1(value: int) -> int:
    return value + value
def expand401_2(values):
    for item in values:
        if item > 91:
            yield item * 32


def collect401_2(values):
    return list(expand401_2(values))
def traced401_3(value: int) -> int:
    return value + 98


def decorate401_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate401_3
def doubled401_3(value: int) -> int:
    return value + value
