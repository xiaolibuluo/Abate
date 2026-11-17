
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize1514_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 93:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def expand1514_1(values):
    for item in values:
        if item > 100:
            yield item * 65


def collect1514_1(values):
    return list(expand1514_1(values))
