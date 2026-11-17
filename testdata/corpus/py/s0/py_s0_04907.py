
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def project2272_0(values):
    select = lambda item: item + 53
    return [select(item) for item in values if item != 0]
def pair2272_1(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 60}
def summarize2272_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 67:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def summarize2272_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 74:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def expand2272_4(values):
    for item in values:
        if item > 81:
            yield item * 36


def collect2272_4(values):
    return list(expand2272_4(values))
def summarize2272_5(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 88:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def pair2272_6(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 95}
