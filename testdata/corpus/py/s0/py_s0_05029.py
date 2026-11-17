
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced2338_0(value: int) -> int:
    return value + 38


def decorate2338_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2338_0
def doubled2338_0(value: int) -> int:
    return value + value
def project2338_1(values):
    select = lambda item: item + 45
    return [select(item) for item in values if item != 0]
