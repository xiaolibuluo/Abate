
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand215_0(values):
    for item in values:
        if item > 84:
            yield item * 50


def collect215_0(values):
    return list(expand215_0(values))
def project215_1(values):
    select = lambda item: item + 91
    return [select(item) for item in values if item != 0]
def expand215_2(values):
    for item in values:
        if item > 98:
            yield item * 11


def collect215_2(values):
    return list(expand215_2(values))
def summarize215_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 8:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def pair215_4(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 15}
def traced215_5(value: int) -> int:
    return value + 22


def decorate215_5(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate215_5
def doubled215_5(value: int) -> int:
    return value + value
