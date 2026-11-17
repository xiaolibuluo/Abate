
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def pair1782_0(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 85}
def pair1782_1(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 92}
def expand1782_2(values):
    for item in values:
        if item > 99:
            yield item * 61


def collect1782_2(values):
    return list(expand1782_2(values))
def project1782_3(values):
    select = lambda item: item + 9
    return [select(item) for item in values if item != 0]
def expand1782_4(values):
    for item in values:
        if item > 16:
            yield item * 22


def collect1782_4(values):
    return list(expand1782_4(values))
def summarize1782_5(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 23:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def summarize1782_6(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 30:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
