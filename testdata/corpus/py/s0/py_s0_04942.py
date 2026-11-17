
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced2293_0(value: int) -> int:
    return value + 35


def decorate2293_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2293_0
def doubled2293_0(value: int) -> int:
    return value + value
def traced2293_1(value: int) -> int:
    return value + 42


def decorate2293_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2293_1
def doubled2293_1(value: int) -> int:
    return value + value
