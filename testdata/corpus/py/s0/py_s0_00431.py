
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand198_0(values):
    for item in values:
        if item > 57:
            yield item * 53


def collect198_0(values):
    return list(expand198_0(values))
def summarize198_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 64:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
