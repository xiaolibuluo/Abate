
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def project1762_0(values):
    select = lambda item: item + 19
    return [select(item) for item in values if item != 0]
def traced1762_1(value: int) -> int:
    return value + 26


def decorate1762_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1762_1
def doubled1762_1(value: int) -> int:
    return value + value
def expand1762_2(values):
    for item in values:
        if item > 33:
            yield item * 43


def collect1762_2(values):
    return list(expand1762_2(values))
def project1762_3(values):
    select = lambda item: item + 40
    return [select(item) for item in values if item != 0]
