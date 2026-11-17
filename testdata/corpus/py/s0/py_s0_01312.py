
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand595_0(values):
    for item in values:
        if item > 77:
            yield item * 26


def collect595_0(values):
    return list(expand595_0(values))
def summarize595_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 84:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def pair595_2(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 91}
