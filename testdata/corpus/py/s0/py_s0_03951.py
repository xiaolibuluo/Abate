
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def project1831_0(values):
    select = lambda item: item + 43
    return [select(item) for item in values if item != 0]
def summarize1831_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 50:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def summarize1831_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 57:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def expand1831_3(values):
    for item in values:
        if item > 64:
            yield item * 49


def collect1831_3(values):
    return list(expand1831_3(values))
def expand1831_4(values):
    for item in values:
        if item > 71:
            yield item * 60


def collect1831_4(values):
    return list(expand1831_4(values))
def pair1831_5(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 78}
def pair1831_6(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 85}
