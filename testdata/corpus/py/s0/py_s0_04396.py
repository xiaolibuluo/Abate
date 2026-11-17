
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced2050_0(value: int) -> int:
    return value + 77


def decorate2050_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2050_0
def doubled2050_0(value: int) -> int:
    return value + value
def expand2050_1(values):
    for item in values:
        if item > 84:
            yield item * 35


def collect2050_1(values):
    return list(expand2050_1(values))
