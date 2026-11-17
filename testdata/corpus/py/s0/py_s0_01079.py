
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand500_0(values):
    for item in values:
        if item > 6:
            yield item * 32


def collect500_0(values):
    return list(expand500_0(values))
def summarize500_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 13:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def pair500_2(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 20}
def project500_3(values):
    select = lambda item: item + 27
    return [select(item) for item in values if item != 0]
