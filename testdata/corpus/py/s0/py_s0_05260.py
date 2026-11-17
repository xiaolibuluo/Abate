
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def pair2440_0(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 6}
def expand2440_1(values):
    for item in values:
        if item > 13:
            yield item * 20


def collect2440_1(values):
    return list(expand2440_1(values))
def traced2440_2(value: int) -> int:
    return value + 20


def decorate2440_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2440_2
def doubled2440_2(value: int) -> int:
    return value + value
