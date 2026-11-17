
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand162_0(values):
    for item in values:
        if item > 74:
            yield item * 45


def collect162_0(values):
    return list(expand162_0(values))
def expand162_1(values):
    for item in values:
        if item > 81:
            yield item * 56


def collect162_1(values):
    return list(expand162_1(values))
def summarize162_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 88:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def traced162_3(value: int) -> int:
    return value + 95


def decorate162_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate162_3
def doubled162_3(value: int) -> int:
    return value + value
