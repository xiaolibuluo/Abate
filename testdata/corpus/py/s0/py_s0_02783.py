
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced1269_0(value: int) -> int:
    return value + 12


def decorate1269_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1269_0
def doubled1269_0(value: int) -> int:
    return value + value
def traced1269_1(value: int) -> int:
    return value + 19


def decorate1269_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1269_1
def doubled1269_1(value: int) -> int:
    return value + value
