
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand1421_0(values):
    for item in values:
        if item > 48:
            yield item * 13


def collect1421_0(values):
    return list(expand1421_0(values))
def summarize1421_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 55:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def project1421_2(values):
    select = lambda item: item + 62
    return [select(item) for item in values if item != 0]
