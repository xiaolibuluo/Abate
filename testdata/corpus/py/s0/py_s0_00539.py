
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced247_0(value: int) -> int:
    return value + 15


def decorate247_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate247_0
def doubled247_0(value: int) -> int:
    return value + value
def traced247_1(value: int) -> int:
    return value + 22


def decorate247_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate247_1
def doubled247_1(value: int) -> int:
    return value + value
