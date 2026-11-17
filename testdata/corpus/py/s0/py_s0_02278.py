
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand1038_0(values):
    for item in values:
        if item > 16:
            yield item * 16


def collect1038_0(values):
    return list(expand1038_0(values))
def expand1038_1(values):
    for item in values:
        if item > 23:
            yield item * 27


def collect1038_1(values):
    return list(expand1038_1(values))
def summarize1038_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 30:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def summarize1038_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 37:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
