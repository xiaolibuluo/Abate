
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize592_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 38:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def expand592_1(values):
    for item in values:
        if item > 45:
            yield item * 16


def collect592_1(values):
    return list(expand592_1(values))
def project592_2(values):
    select = lambda item: item + 52
    return [select(item) for item in values if item != 0]
