
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced1658_0(value: int) -> int:
    return value + 25


def decorate1658_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1658_0
def doubled1658_0(value: int) -> int:
    return value + value
def project1658_1(values):
    select = lambda item: item + 32
    return [select(item) for item in values if item != 0]
