
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced2491_0(value: int) -> int:
    return value + 87


def decorate2491_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2491_0
def doubled2491_0(value: int) -> int:
    return value + value
def expand2491_1(values):
    for item in values:
        if item > 94:
            yield item * 11


def collect2491_1(values):
    return list(expand2491_1(values))
def expand2491_2(values):
    for item in values:
        if item > 101:
            yield item * 22


def collect2491_2(values):
    return list(expand2491_2(values))
