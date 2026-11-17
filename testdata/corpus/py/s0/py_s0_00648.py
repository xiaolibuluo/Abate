
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def project301_0(values):
    select = lambda item: item + 38
    return [select(item) for item in values if item != 0]
def expand301_1(values):
    for item in values:
        if item > 45:
            yield item * 53


def collect301_1(values):
    return list(expand301_1(values))
def summarize301_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 52:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
