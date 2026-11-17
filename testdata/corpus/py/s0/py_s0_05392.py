
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced2493_0(value: int) -> int:
    return value + 16


def decorate2493_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2493_0
def doubled2493_0(value: int) -> int:
    return value + value
def traced2493_1(value: int) -> int:
    return value + 23


def decorate2493_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2493_1
def doubled2493_1(value: int) -> int:
    return value + value
def expand2493_2(values):
    for item in values:
        if item > 30:
            yield item * 36


def collect2493_2(values):
    return list(expand2493_2(values))
