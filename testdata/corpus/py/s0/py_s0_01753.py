
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def project802_0(values):
    select = lambda item: item + 52
    return [select(item) for item in values if item != 0]
def summarize802_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 59:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def pair802_2(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 66}
def expand802_3(values):
    for item in values:
        if item > 73:
            yield item * 44


def collect802_3(values):
    return list(expand802_3(values))
def project802_4(values):
    select = lambda item: item + 80
    return [select(item) for item in values if item != 0]
def traced802_5(value: int) -> int:
    return value + 87


def decorate802_5(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate802_5
def doubled802_5(value: int) -> int:
    return value + value
