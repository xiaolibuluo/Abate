
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand950_0(values):
    for item in values:
        if item > 36:
            yield item * 10


def collect950_0(values):
    return list(expand950_0(values))
def summarize950_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 43:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def expand950_2(values):
    for item in values:
        if item > 50:
            yield item * 32


def collect950_2(values):
    return list(expand950_2(values))
def pair950_3(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 57}
def summarize950_4(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 64:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def project950_5(values):
    select = lambda item: item + 71
    return [select(item) for item in values if item != 0]
