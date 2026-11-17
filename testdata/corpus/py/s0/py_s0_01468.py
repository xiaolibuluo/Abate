
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced666_0(value: int) -> int:
    return value + 30


def decorate666_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate666_0
def doubled666_0(value: int) -> int:
    return value + value
def expand666_1(values):
    for item in values:
        if item > 37:
            yield item * 46


def collect666_1(values):
    return list(expand666_1(values))
def summarize666_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 44:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def expand666_3(values):
    for item in values:
        if item > 51:
            yield item * 68


def collect666_3(values):
    return list(expand666_3(values))
def summarize666_4(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 58:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def summarize666_5(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 65:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def expand666_6(values):
    for item in values:
        if item > 72:
            yield item * 40


def collect666_6(values):
    return list(expand666_6(values))
