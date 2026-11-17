
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced295_0(value: int) -> int:
    return value + 57


def decorate295_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate295_0
def doubled295_0(value: int) -> int:
    return value + value
def project295_1(values):
    select = lambda item: item + 64
    return [select(item) for item in values if item != 0]
