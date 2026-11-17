
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced299_0(value: int) -> int:
    return value + 12


def decorate299_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate299_0
def doubled299_0(value: int) -> int:
    return value + value
def project299_1(values):
    select = lambda item: item + 19
    return [select(item) for item in values if item != 0]
def expand299_2(values):
    for item in values:
        if item > 26:
            yield item * 50


def collect299_2(values):
    return list(expand299_2(values))
