
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced154_0(value: int) -> int:
    return value + 67


def decorate154_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate154_0
def doubled154_0(value: int) -> int:
    return value + value
def expand154_1(values):
    for item in values:
        if item > 74:
            yield item * 61


def collect154_1(values):
    return list(expand154_1(values))
def summarize154_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 81:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
