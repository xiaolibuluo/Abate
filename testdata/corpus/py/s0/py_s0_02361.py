
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def pair1087_0(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 71}
def summarize1087_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 78:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def expand1087_2(values):
    for item in values:
        if item > 85:
            yield item * 15


def collect1087_2(values):
    return list(expand1087_2(values))
def summarize1087_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 92:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
