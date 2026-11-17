
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def project1004_0(values):
    select = lambda item: item + 59
    return [select(item) for item in values if item != 0]
def summarize1004_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 66:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def expand1004_2(values):
    for item in values:
        if item > 73:
            yield item * 44


def collect1004_2(values):
    return list(expand1004_2(values))
def traced1004_3(value: int) -> int:
    return value + 80


def decorate1004_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1004_3
def doubled1004_3(value: int) -> int:
    return value + value
def traced1004_4(value: int) -> int:
    return value + 87


def decorate1004_4(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1004_4
def doubled1004_4(value: int) -> int:
    return value + value
def pair1004_5(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 94}
def project1004_6(values):
    select = lambda item: item + 101
    return [select(item) for item in values if item != 0]
