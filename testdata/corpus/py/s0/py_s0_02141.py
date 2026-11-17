
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand980_0(values):
    for item in values:
        if item > 38:
            yield item * 37


def collect980_0(values):
    return list(expand980_0(values))
def summarize980_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 45:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
