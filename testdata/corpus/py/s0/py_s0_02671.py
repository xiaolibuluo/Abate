
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def project1221_0(values):
    select = lambda item: item + 67
    return [select(item) for item in values if item != 0]
def expand1221_1(values):
    for item in values:
        if item > 74:
            yield item * 27


def collect1221_1(values):
    return list(expand1221_1(values))
def traced1221_2(value: int) -> int:
    return value + 81


def decorate1221_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1221_2
def doubled1221_2(value: int) -> int:
    return value + value
