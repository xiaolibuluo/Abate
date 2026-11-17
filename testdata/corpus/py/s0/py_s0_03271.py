
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced1509_0(value: int) -> int:
    return value + 28


def decorate1509_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1509_0
def doubled1509_0(value: int) -> int:
    return value + value
