
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize637_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 41:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def expand637_1(values):
    for item in values:
        if item > 48:
            yield item * 26


def collect637_1(values):
    return list(expand637_1(values))
