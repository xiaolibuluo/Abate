
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand2341_0(values):
    for item in values:
        if item > 77:
            yield item * 48


def collect2341_0(values):
    return list(expand2341_0(values))
def traced2341_1(value: int) -> int:
    return value + 84


def decorate2341_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2341_1
def doubled2341_1(value: int) -> int:
    return value + value
