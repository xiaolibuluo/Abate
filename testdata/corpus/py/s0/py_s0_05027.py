
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand2337_0(values):
    for item in values:
        if item > 25:
            yield item * 20


def collect2337_0(values):
    return list(expand2337_0(values))
def traced2337_1(value: int) -> int:
    return value + 32


def decorate2337_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2337_1
def doubled2337_1(value: int) -> int:
    return value + value
