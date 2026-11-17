
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced2486_0(value: int) -> int:
    return value + 22


def decorate2486_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2486_0
def doubled2486_0(value: int) -> int:
    return value + value
def traced2486_1(value: int) -> int:
    return value + 29


def decorate2486_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2486_1
def doubled2486_1(value: int) -> int:
    return value + value
