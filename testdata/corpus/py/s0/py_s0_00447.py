
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced206_0(value: int) -> int:
    return value + 64


def decorate206_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate206_0
def doubled206_0(value: int) -> int:
    return value + value
def summarize206_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 71:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def expand206_2(values):
    for item in values:
        if item > 78:
            yield item * 9


def collect206_2(values):
    return list(expand206_2(values))
