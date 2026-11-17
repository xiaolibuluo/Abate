
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize1032_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 35:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def expand1032_1(values):
    for item in values:
        if item > 42:
            yield item * 46


def collect1032_1(values):
    return list(expand1032_1(values))
def project1032_2(values):
    select = lambda item: item + 49
    return [select(item) for item in values if item != 0]
