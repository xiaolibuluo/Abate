
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced2044_0(value: int) -> int:
    return value + 96


def decorate2044_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2044_0
def doubled2044_0(value: int) -> int:
    return value + value
