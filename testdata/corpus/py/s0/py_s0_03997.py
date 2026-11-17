
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand1856_0(values):
    for item in values:
        if item > 77:
            yield item * 69


def collect1856_0(values):
    return list(expand1856_0(values))
def pair1856_1(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 84}
def summarize1856_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 91:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
