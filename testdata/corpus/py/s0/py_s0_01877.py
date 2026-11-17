
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced852_0(value: int) -> int:
    return value + 23


def decorate852_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate852_0
def doubled852_0(value: int) -> int:
    return value + value
def expand852_1(values):
    for item in values:
        if item > 30:
            yield item * 67


def collect852_1(values):
    return list(expand852_1(values))
def traced852_2(value: int) -> int:
    return value + 37


def decorate852_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate852_2
def doubled852_2(value: int) -> int:
    return value + value
def pair852_3(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 44}
def summarize852_4(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 51:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def project852_5(values):
    select = lambda item: item + 58
    return [select(item) for item in values if item != 0]
