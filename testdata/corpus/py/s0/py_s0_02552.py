
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize1177_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 77:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def expand1177_1(values):
    for item in values:
        if item > 84:
            yield item * 24


def collect1177_1(values):
    return list(expand1177_1(values))
def project1177_2(values):
    select = lambda item: item + 91
    return [select(item) for item in values if item != 0]
