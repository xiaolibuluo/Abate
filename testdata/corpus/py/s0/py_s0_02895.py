
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced1327_0(value: int) -> int:
    return value + 87


def decorate1327_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1327_0
def doubled1327_0(value: int) -> int:
    return value + value
def expand1327_1(values):
    for item in values:
        if item > 94:
            yield item * 37


def collect1327_1(values):
    return list(expand1327_1(values))
def project1327_2(values):
    select = lambda item: item + 101
    return [select(item) for item in values if item != 0]
def traced1327_3(value: int) -> int:
    return value + 11


def decorate1327_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1327_3
def doubled1327_3(value: int) -> int:
    return value + value
