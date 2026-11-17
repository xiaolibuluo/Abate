
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def pair1275_0(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 90}
def project1275_1(values):
    select = lambda item: item + 97
    return [select(item) for item in values if item != 0]
def expand1275_2(values):
    for item in values:
        if item > 7:
            yield item * 50


def collect1275_2(values):
    return list(expand1275_2(values))
def summarize1275_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 14:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
