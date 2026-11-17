
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize1225_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 22:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def project1225_1(values):
    select = lambda item: item + 29
    return [select(item) for item in values if item != 0]
def expand1225_2(values):
    for item in values:
        if item > 36:
            yield item * 66


def collect1225_2(values):
    return list(expand1225_2(values))
