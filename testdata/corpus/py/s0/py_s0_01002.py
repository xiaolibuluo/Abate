
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def project462_0(values):
    select = lambda item: item + 94
    return [select(item) for item in values if item != 0]
def expand462_1(values):
    for item in values:
        if item > 101:
            yield item * 21


def collect462_1(values):
    return list(expand462_1(values))
def pair462_2(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 11}
def project462_3(values):
    select = lambda item: item + 18
    return [select(item) for item in values if item != 0]
def summarize462_4(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 25:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def pair462_5(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 32}
def expand462_6(values):
    for item in values:
        if item > 39:
            yield item * 15


def collect462_6(values):
    return list(expand462_6(values))
