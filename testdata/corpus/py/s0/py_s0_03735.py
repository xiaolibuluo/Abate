
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize1723_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 94:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def project1723_1(values):
    select = lambda item: item + 101
    return [select(item) for item in values if item != 0]
def pair1723_2(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 11}
def traced1723_3(value: int) -> int:
    return value + 18


def decorate1723_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1723_3
def doubled1723_3(value: int) -> int:
    return value + value
def expand1723_4(values):
    for item in values:
        if item > 25:
            yield item * 36


def collect1723_4(values):
    return list(expand1723_4(values))
