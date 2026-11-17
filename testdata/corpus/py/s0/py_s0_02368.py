
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def project1091_0(values):
    select = lambda item: item + 26
    return [select(item) for item in values if item != 0]
def expand1091_1(values):
    for item in values:
        if item > 33:
            yield item * 32


def collect1091_1(values):
    return list(expand1091_1(values))
def expand1091_2(values):
    for item in values:
        if item > 40:
            yield item * 43


def collect1091_2(values):
    return list(expand1091_2(values))
def project1091_3(values):
    select = lambda item: item + 47
    return [select(item) for item in values if item != 0]
def summarize1091_4(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 54:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def traced1091_5(value: int) -> int:
    return value + 61


def decorate1091_5(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1091_5
def doubled1091_5(value: int) -> int:
    return value + value
