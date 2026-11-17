
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize644_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 35:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def expand644_1(values):
    for item in values:
        if item > 42:
            yield item * 14


def collect644_1(values):
    return list(expand644_1(values))
def traced644_2(value: int) -> int:
    return value + 49


def decorate644_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate644_2
def doubled644_2(value: int) -> int:
    return value + value
