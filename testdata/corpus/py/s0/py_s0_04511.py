
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced2102_0(value: int) -> int:
    return value + 74


def decorate2102_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2102_0
def doubled2102_0(value: int) -> int:
    return value + value
def expand2102_1(values):
    for item in values:
        if item > 81:
            yield item * 33


def collect2102_1(values):
    return list(expand2102_1(values))
def expand2102_2(values):
    for item in values:
        if item > 88:
            yield item * 44


def collect2102_2(values):
    return list(expand2102_2(values))
