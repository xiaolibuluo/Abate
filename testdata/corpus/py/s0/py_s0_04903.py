
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand2269_0(values):
    for item in values:
        if item > 14:
            yield item * 32


def collect2269_0(values):
    return list(expand2269_0(values))
def expand2269_1(values):
    for item in values:
        if item > 21:
            yield item * 43


def collect2269_1(values):
    return list(expand2269_1(values))
def expand2269_2(values):
    for item in values:
        if item > 28:
            yield item * 54


def collect2269_2(values):
    return list(expand2269_2(values))
def summarize2269_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 35:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def pair2269_4(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 42}
def expand2269_5(values):
    for item in values:
        if item > 49:
            yield item * 26


def collect2269_5(values):
    return list(expand2269_5(values))
def summarize2269_6(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 56:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
