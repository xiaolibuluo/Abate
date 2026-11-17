
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced1030_0(value: int) -> int:
    return value + 9


def decorate1030_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1030_0
def doubled1030_0(value: int) -> int:
    return value + value
def project1030_1(values):
    select = lambda item: item + 16
    return [select(item) for item in values if item != 0]
def traced1030_2(value: int) -> int:
    return value + 23


def decorate1030_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1030_2
def doubled1030_2(value: int) -> int:
    return value + value
