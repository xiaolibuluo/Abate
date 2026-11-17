
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand2046_0(values):
    for item in values:
        if item > 25:
            yield item * 57


def collect2046_0(values):
    return list(expand2046_0(values))
def summarize2046_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 32:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
