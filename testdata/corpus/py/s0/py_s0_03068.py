
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand1414_0(values):
    for item in values:
        if item > 54:
            yield item * 25


def collect1414_0(values):
    return list(expand1414_0(values))
def summarize1414_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 61:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
