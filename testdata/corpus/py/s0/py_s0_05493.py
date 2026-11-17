
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced2537_0(value: int) -> int:
    return value + 6


def decorate2537_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2537_0
def doubled2537_0(value: int) -> int:
    return value + value
def project2537_1(values):
    select = lambda item: item + 13
    return [select(item) for item in values if item != 0]
def project2537_2(values):
    select = lambda item: item + 20
    return [select(item) for item in values if item != 0]
